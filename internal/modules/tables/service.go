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
	return this.repo.ListAllTables(ctx)
}

func (this *TableService) FindTableById(ctx context.Context, id string) (DBTables, error) {
	return this.repo.FindTableById(ctx, id)
}

func (this *TableService) IsAvailable(ctx context.Context, id string) (bool, error) {
	return this.repo.IsAvailable(ctx, id)
}
