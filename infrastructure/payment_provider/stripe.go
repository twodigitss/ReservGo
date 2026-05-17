package paymentprovider

import (
	"context"
	"fmt"
	"time"
	"github.com/twodigitss/reserv-go/internal/modules/payment"
)

// var _ payment.Provider = (*StripeImpl)(nil)
type StripeImpl struct { }
func NewStripe() *StripeImpl{
	return &StripeImpl{}
}

func (this *StripeImpl) Pay(ctx context.Context, _body payment.DBPayment) (payment.ProviderTransaction,error){
	fmt.Println("Recieved $1 %2 from $3!", _body.Amount, _body.Currency, _body.ClientUUID)

	//TODO: would be great to send the body through email
  res := payment.ProviderTransaction{
		TransId: "pi_1NXYZ1234567890abcdef",
		CreatedAt: time.Now(),
		Confirmation: true,
		TotalAmount: _body.Amount,
		Currency: _body.Currency,
		SenderAcc: _body.FromAcc,
		RecieverAcc: "cus_ABC1234567892",
  }

	return res, nil
}

func (this *StripeImpl) Refund(ctx context.Context, _body payment.DBPayment) (payment.ProviderTransaction,error){
	fmt.Println("Processing refund of $1 for $2. check your email.", _body.ClientUUID, _body.Amount)

  res := payment.ProviderTransaction{
		TransId: _body.UUID,
		CreatedAt: _body.CreatedAt,
		Confirmation: true,
		TotalAmount: _body.Amount,
		Currency: _body.Currency,
		SenderAcc: _body.FromAcc,
		RecieverAcc: "cus_ABC1234567892",
  }

	return res, nil

}
