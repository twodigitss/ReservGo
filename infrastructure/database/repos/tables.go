package repos

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twodigitss/reserv-go/internal/modules/tables"
)

// var is_this_file_implemmenting_correclty_the_interface tables.TableModuleInterface = (*TableRepoImpl)(nil)

type TableRepoImpl struct { DB *pgxpool.Pool }

func NewTableRepo(db *pgxpool.Pool) *TableRepoImpl{
	return &TableRepoImpl{DB: db}
}

//Returns all the table "Tables". This is a bad (really really bad ) practice irl.
//it is merely for demonstrable purposes
func (this *TableRepoImpl) ListAllTables(ctx context.Context) ([]tables.DBTables, error){
	rows, err := this.DB.Query( ctx, 
		`SELECT id, description, capacity, floor, created_at
		 FROM dbv2.tables`,
	)

	if err != nil { return nil, err }
	defer rows.Close()

	var myRows []tables.DBTables;
	for rows.Next() {
		var thing tables.DBTables

		err := rows.Scan(&thing.Id, &thing.Description, &thing.Capacity, &thing.Floor, &thing.CreatedAt)
		if err != nil {
			fmt.Print("scan failed: %w", err)
			return nil, err
		}

		myRows = append(myRows, thing)
	}

	if err := rows.Err(); err != nil {
		fmt.Print("rows iteration error: %w", err)
		return nil, err
	}

	return myRows,nil
}

func (this *TableRepoImpl) FindTableById(ctx context.Context, _id string) (tables.DBTables, error){
	var result tables.DBTables
	err := this.DB.QueryRow( ctx, 
		`SELECT id, description, capacity, floor, created_at
		 FROM dbv2.tables 
		 WHERE id = $1`, _id,
	).Scan(&result.Id, &result.Description, &result.Capacity, &result.Floor, &result.CreatedAt)

	if err != nil {
		fmt.Print("ERROR GETTING PGX DATA: ", err)
		return tables.DBTables{}, err
	}
	return result, nil
}


func (this *TableRepoImpl) IsAvailable(ctx context.Context, id string) (bool, error) {
	var result bool
	err := this.DB.QueryRow( ctx, 
		`SELECT NOT EXISTS ( 
			SELECT 1 FROM dbv2.reservations WHERE table_id = $1 AND visited = false
		) 
		AS available;`, id,
	).Scan(result)

	if err != nil {
		return false, err
	}
	return result, nil
}
