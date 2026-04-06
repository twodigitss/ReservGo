package supabase
import (
	"context"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twodigitss/reserv-go/configs"
)

func Connect() (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(
		context.Background(), 
		configs.DB_URL,
	)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	return conn,nil
}
