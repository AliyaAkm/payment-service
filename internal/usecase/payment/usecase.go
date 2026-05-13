package payment

type UseCase struct {
	orderRepo        OrderRepository
	paymentRepo      PaymentRepo
	paymentClient    PaymentClient
	orderStatusRepo  OrderStatusRepo
	subscriptionRepo SubscriptionRepo
	PublicKey        string
	TerminalID       string //  идентификатор магазина
}

func New(orderRepo OrderRepository, paymentRepo PaymentRepo, paymentClient PaymentClient, orderStatusRepo OrderStatusRepo, subscriptionRepo SubscriptionRepo, PublicKey string, TerminalID string) *UseCase {
	return &UseCase{
		orderRepo:        orderRepo,
		paymentRepo:      paymentRepo,
		paymentClient:    paymentClient,
		orderStatusRepo:  orderStatusRepo,
		subscriptionRepo: subscriptionRepo,
		PublicKey:        PublicKey,
		TerminalID:       TerminalID,
	}
}
