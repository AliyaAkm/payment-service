package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"payment-service/internal/http/middleware"
	"payment-service/internal/http/router"
	"syscall"

	orderhandler "payment-service/internal/http/handlers/order"
	pricehandler "payment-service/internal/http/handlers/price"
	orderrepo "payment-service/internal/repo/postgres/order"
	orderstatusrepo "payment-service/internal/repo/postgres/orderstatus"
	pricerepo "payment-service/internal/repo/postgres/price"
	orderusecase "payment-service/internal/usecase/order"
	priceusecase "payment-service/internal/usecase/price"
)

func main() {
	err := godotenv.Load(".env")

	cfg, err := ReadEnv()
	if err != nil {
		log.Fatal("configuration error:", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	pool, err := NewPool(
		ctx,
		cfg.DB,
	)
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

	handler := router.Handler{
		Price: priceHandler,
		Order: orderHandler,
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
