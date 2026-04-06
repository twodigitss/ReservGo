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

func (this *TableService) ListAllTables(ctx context.Context) (*[]DBTables, error) {
	// delega al repo, no reimplementa
	// perfectamente puedo poner aqui logica que maneje el ouput 
	return this.repo.ListAllTables(ctx)
}
