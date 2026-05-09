package reservation
import "context"

type ReservService struct {
    repo Reservation
}
func NewService(repo Reservation) *ReservService {
    return &ReservService{repo: repo}
}

func (this *ReservService) Book(ctx context.Context, body DBReservation) error{
	err := this.repo.Book(ctx, body)
	if err != nil { return err }
	return nil
}

func (this *ReservService) Cancel(ctx context.Context, _id string) error{
	err := this.repo.Cancel(ctx, _id)
	if err != nil { return err }
	return nil
}

func (this *ReservService) Update(ctx context.Context, _id string) error{
	err := this.repo.Update(ctx, _id)
	if err != nil { return err }
	return nil
}

func (this *ReservService) Exists(ctx context.Context, _id string) (bool, error){
	bo, err := this.repo.Exists(ctx, _id)
	if err != nil { return false, err }
	return bo, nil
}

func (this *ReservService) GetByID(ctx context.Context, _id string) (*DBReservation, error){
	res, err := this.repo.GetByID(ctx, _id)
	if err != nil { return nil, err }

	return res, nil
}

func (this *ReservService) GetByClientUUID(ctx context.Context, _id string) (*[]DBReservation, error){
	res, err := this.repo.GetByClientUUID(ctx, _id)
	if err != nil { return nil, err }

	return res, nil
}
