package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"payment-service/internal/http/middleware"
	"payment-service/internal/http/router"
	"payment-service/internal/service/notification"
	"syscall"

	epaymenthandler "payment-service/internal/http/handlers/epayment"
	orderhandler "payment-service/internal/http/handlers/order"
	paymentmethodhandler "payment-service/internal/http/handlers/paymentmethod"
	pricehandler "payment-service/internal/http/handlers/price"
	orderrepo "payment-service/internal/repo/postgres/order"
	orderstatusrepo "payment-service/internal/repo/postgres/orderstatus"
	paymentrepo "payment-service/internal/repo/postgres/payment"
	paymentmethodrepo "payment-service/internal/repo/postgres/paymentmethod"
	pricerepo "payment-service/internal/repo/postgres/price"
	subscriptionrepo "payment-service/internal/repo/postgres/subscription"
	orderusecase "payment-service/internal/usecase/order"
	paymentusecase "payment-service/internal/usecase/payment"
	paymentmethodusecase "payment-service/internal/usecase/paymentmethod"
	priceusecase "payment-service/internal/usecase/price"

	paymentclient "payment-service/service/epayment"
)

func main() {
	err := godotenv.Load(".env")

	cfg, err := ReadEnv()
	if err != nil {
		log.Fatal("configuration error:", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("error to log", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	pool, err := NewPool(
		ctx,
		cfg.DB,
	)

	payment := paymentclient.NewClient(logger, cfg.Epayment.Duration, cfg.Epayment.AuthHost, cfg.Epayment.Host, cfg.Epayment.TerminalID, cfg.Epayment.ClientID, cfg.Epayment.ClientSecret, cfg.Epayment.SecretHash)
	if err != nil {
		log.Fatal("error connecting to the database pool:", err)
	}
	defer pool.Close()

	db, err := NewDB(ctx, cfg.DB)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	// order status
	orderstatusRepo := orderstatusrepo.NewRepo(db)

	priceRepo := pricerepo.New(db)
	priceUseCase := priceusecase.New(priceRepo)
	priceHandler := pricehandler.New(priceUseCase)

	// order
	orderRepo := orderrepo.NewRepo(db)
	orderUseCase := orderusecase.New(orderRepo, orderstatusRepo, priceRepo)
	orderHandler := orderhandler.New(orderUseCase)

	// payment method
	paymentMethodRepo := paymentmethodrepo.NewRepo(db)
	paymentMethodUseCase := paymentmethodusecase.New(paymentMethodRepo)
	paymentMethodHandler := paymentmethodhandler.NewHandler(paymentMethodUseCase)

	// subscription
	subscriptionRepo := subscriptionrepo.New(db)

	var notificationClient paymentusecase.NotificationSender
	if cfg.Notification.InternalAPIKey != "" {
		notificationClient, err = notification.NewClient(notification.ClientConfig{
			BaseURL:        cfg.Notification.URL,
			Timeout:        cfg.Notification.Timeout,
			InternalAPIKey: cfg.Notification.InternalAPIKey,
		})
		if err != nil {
			log.Fatal("error configuring notification service client:", err)
		}
	}

	// payment
	paymentRepo := paymentrepo.NewRepo(db)
	paymentUseCase := paymentusecase.New(orderRepo, paymentRepo, payment, orderstatusRepo, subscriptionRepo, cfg.Epayment.PublicKey, cfg.Epayment.TerminalID, notificationClient)
	paymentHandler := epaymenthandler.NewHandler(payment, paymentUseCase)

	handler := router.Handler{
		Price:         priceHandler,
		Order:         orderHandler,
		PaymentMethod: paymentMethodHandler,
		Payment:       paymentHandler,
	}

	engine := router.New(
		handler,
		[]gin.HandlerFunc{
			middleware.RequestID(),
			middleware.Logger(),
			middleware.Recover(),
		},
	)

	srv := &http.Server{
		Addr:              cfg.ListenAddr(),
		Handler:           engine,
		ReadHeaderTimeout: cfg.HTTP.ReadHeaderTimeout,
	}

	go func() {
		log.Println("curriculum-service started on", cfg.ListenAddr())
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("server error:", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.HTTP.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Println("shutdown error:", err)
	}

	log.Println("curriculum-service stopped")
}
