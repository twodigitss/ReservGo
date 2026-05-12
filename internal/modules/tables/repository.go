package tables
//aqui declaro los contratos que deberia de tener infra/supabase/repos/tables.go

import "context"

//recordar que en mittal me ensenaron que es bueno tambien poner
//el nombre del storedProc como nombre del metodo porque asi te
//confundes menos

type TableModuleInterface interface {
	ListAllTables(ctx context.Context) ([]DBTables, error)
	FindTableById(ctx context.Context, uuid string) (DBTables, error)
	SetTableAvailable(ctx context.Context, uuid string) (DBTables, error)
	SetTableOccupied(ctx context.Context, uuid string) (DBTables, error)
}
