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

func (this *ReservService) Cancel(ctx context.Context, _uuid string) error{
	return this.repo.Cancel(ctx, _uuid)
}

func (this *ReservService) Update(ctx context.Context, _uuid string) error{
	return this.repo.Update(ctx, _uuid)
}

func (this *ReservService) Exists(ctx context.Context, _uuid string) (bool, error){
	return this.repo.Exists(ctx, _uuid)
}

func (this *ReservService) GetByID(ctx context.Context, _uuid string) (DBReservation, error){
	return this.repo.GetByID(ctx, _uuid)
}

func (this *ReservService) GetByClientUUID(ctx context.Context, _uuid string) ([]DBReservation, error){
	return this.repo.GetByClientUUID(ctx, _uuid)
}
