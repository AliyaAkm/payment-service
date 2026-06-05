package payment

type UseCase struct {
	orderRepo        OrderRepository
	paymentRepo      PaymentRepo
	paymentClient    PaymentClient
	orderStatusRepo  OrderStatusRepo
	subscriptionRepo SubscriptionRepo
	notification     NotificationSender
	PublicKey        string
	TerminalID       string //  идентификатор магазина
}

func New(orderRepo OrderRepository, paymentRepo PaymentRepo, paymentClient PaymentClient, orderStatusRepo OrderStatusRepo, subscriptionRepo SubscriptionRepo, PublicKey string, TerminalID string, notification ...NotificationSender) *UseCase {
	var sender NotificationSender
	if len(notification) > 0 {
		sender = notification[0]
	}
	return &UseCase{
		orderRepo:        orderRepo,
		paymentRepo:      paymentRepo,
		paymentClient:    paymentClient,
		orderStatusRepo:  orderStatusRepo,
		subscriptionRepo: subscriptionRepo,
		notification:     sender,
		PublicKey:        PublicKey,
		TerminalID:       TerminalID,
	}
}
