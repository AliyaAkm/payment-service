package epayment

type Handler struct {
	client         client
	paymentUseCase paymentUseCase
}

func NewHandler(client client, paymentUseCase paymentUseCase) *Handler {
	return &Handler{
		client:         client,
		paymentUseCase: paymentUseCase,
	}
}
