package payment
import "context"

type Payment interface {
	SaveToDB(ctx context.Context, _body DBPayment)(error)
	GetByTransID(ctx context.Context, _id string)(error)
}

type Provider interface {
	Pay(ctx context.Context, _body DBPayment) (ProviderTransaction, error)
	Refund(ctx context.Context, _body DBPayment) (ProviderTransaction, error)
}
