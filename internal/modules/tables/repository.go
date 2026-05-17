package tables
//aqui declaro los contratos que deberia de tener infra/supabase/repos/tables.go

import "context"

type TableModuleInterface interface {
	ListAllTables(ctx context.Context) ([]DBTables, error)
	FindTableById(ctx context.Context, id string) (DBTables, error)
	IsAvailable(ctx context.Context, id string) (bool, error)
}
