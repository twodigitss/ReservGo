package tables
import (
	"context"
)

type TableService struct {
    repo TableModuleInterface
}

func NewService(repo TableModuleInterface) *TableService {
    return &TableService{repo: repo}
}

func (this *TableService) ListAllTables(ctx context.Context) ([]DBTables, error) {
	// delega al repo, no reimplementa
	// perfectamente puedo poner aqui logica que maneje el ouput
	return this.repo.ListAllTables(ctx)
}

func (this *TableService) FindTableById(ctx context.Context, uuid string) (*DBTables, error) {
	return this.repo.FindTableById(ctx, uuid)
}
func (this *TableService) SetTableAvailable(ctx context.Context, uuid string) (*DBTables, error) {
	return this.repo.SetTableAvailable(ctx, uuid)
}

func (this *TableService) SetTableOccupied(ctx context.Context, uuid string) (*DBTables, error) {
	return this.repo.SetTableOccupied(ctx, uuid)
}
