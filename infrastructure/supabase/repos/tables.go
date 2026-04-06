package repos

import (
	"context"
	"fmt"
	"time"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twodigitss/reserv-go/internal/modules/tables"
)

// var is_this_file_implemmenting_correclty_the_interface tables.TableModuleInterface = (*TableRepoImpl)(nil)

type TableRepoImpl struct { DB *pgxpool.Pool }

func NewTableRepo(db *pgxpool.Pool) *TableRepoImpl{
	return &TableRepoImpl{DB: db}
}

//Returns the first result of the table "Tables"
func (this *TableRepoImpl) ListFirstTable(ctx context.Context) (*tables.DBTables, error){
	var result tables.DBTables
	err := this.DB.QueryRow(
		context.Background(),
		"select id,created_at,reserved,table_name from reservations_demo.dinner_tables",
	).Scan(&result.Id, &result.CreatedAt, &result.Reserved, &result.TableName)

	if err != nil {
		fmt.Print("ERROR GETTING PGX DATA: ", err)
		return nil, err
	}
	return &result, nil

}

//Returns all the table "Tables"
func (this *TableRepoImpl) ListAllTables(ctx context.Context) (*[]tables.DBTables, error){
	rows, err := this.DB.Query(
		context.Background(), 
		"select id,created_at,reserved,table_name from reservations_demo.dinner_tables",
	)

	if err != nil {
		fmt.Print("ERROR GETTING PGX DATA: ", err)
		return nil, err
	}

	defer rows.Close()

	var myRows []tables.DBTables;

	for rows.Next() {
		var id, table_name string
		var created_at time.Time
		var reserved bool

		err := rows.Scan(&id, &created_at, &reserved, &table_name) 
		if err != nil {
			fmt.Print("scan failed: %w", err)
			return nil, err
		}

		myRows = append(myRows, 
			tables.DBTables{Id: id, CreatedAt: created_at, Reserved: reserved, TableName: table_name},
		)

	}

	if err := rows.Err(); err != nil {
		fmt.Print("rows iteration error: %w", err)
		return nil, err
	}

	return &myRows,nil
}
