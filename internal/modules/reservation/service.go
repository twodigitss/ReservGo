package reservation
import "context"

type ReservService struct {
    repo Reservation
}
func NewService(repo Reservation) *ReservService {
    return &ReservService{repo: repo}
}

func (this *ReservService) Book(ctx context.Context, body DBReservation) (string,error){
	return this.repo.Book(ctx, body)
}

func (this *ReservService) Cancel(ctx context.Context, _id string) error{
	return this.repo.Cancel(ctx, _id)
}

func (this *ReservService) Update(ctx context.Context, _id string) error{
	return this.repo.Update(ctx, _id)
}

func (this *ReservService) Exists(ctx context.Context, _id string) (bool, error){
	return this.repo.Exists(ctx, _id)
}

func (this *ReservService) GetByID(ctx context.Context, _id string) (DBReservation, error){
	return this.repo.GetByID(ctx, _id)
}

func (this *ReservService) GetByClientUUID(ctx context.Context, _id string) ([]DBReservation, error){
	return this.repo.GetByClientUUID(ctx, _id)
}
