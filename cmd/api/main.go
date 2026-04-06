package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/twodigitss/reserv-go/cmd/api/handlers"
	"github.com/twodigitss/reserv-go/configs"
	"github.com/twodigitss/reserv-go/infrastructure/supabase"
	"github.com/twodigitss/reserv-go/infrastructure/supabase/repos"
	"github.com/twodigitss/reserv-go/internal/modules/tables"
	"github.com/twodigitss/reserv-go/internal/modules/users"
)

type Container struct {
    Tables *handlers.TableHandler
    Users *handlers.UserHandler
}

func BuildContainer(pool *pgxpool.Pool) *Container {
		tableService := tables.NewService(repos.NewTableRepo(pool))
		userService := users.NewService(repos.NewUserRepo(pool))

		return &Container{
			Tables: &handlers.TableHandler{Service: tableService},  // handler recibe el service
			Users: &handlers.UserHandler{Service: userService},
		}
}

func main(){
	configs.LoadEnv()
	g := gin.Default()

	conn, err := supabase.Connect()
	if err != nil {
		log.Println("Error connecting to supabase")
	}
	defer conn.Close()

	Routes(g, *BuildContainer(conn))

	g.Run()
}

