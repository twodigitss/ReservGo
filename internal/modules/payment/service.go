package payment
import "context"

type PayService struct { repo Payment }
func NewService(repo Payment) *PayService  {
    return &PayService {repo: repo}
}
	
func (this *PayService) SaveToDB(ctx context.Context, _body DBPayment)(DBPayment, error){
	return this.repo.SaveToDB(ctx, _body)
}

func (this *PayService) GetByTransID(ctx context.Context, _id string)(DBPayment, error){
	return this.repo.GetByTransID(ctx, _id)
}
