package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/twodigitss/reserv-go/cmd/api/handlers"
	"github.com/twodigitss/reserv-go/configs"                       //env vars
	"github.com/twodigitss/reserv-go/infrastructure/supabase"       //conection
	"github.com/twodigitss/reserv-go/infrastructure/supabase/repos" //interface impl.
	"github.com/twodigitss/reserv-go/internal/modules/tables"       //module
	"github.com/twodigitss/reserv-go/internal/modules/users"        //module
)

type Container struct {
	Tables *handlers.TableHandler
	Users *handlers.UserHandler
}

func BuildContainer(pool *pgxpool.Pool) *Container {
	//repos
	tableRepo := repos.NewTableRepo(pool)
	userRepo := repos.NewUserRepo(pool)

	//services
	tableService := tables.NewService(tableRepo)
	userService := users.NewService(userRepo)

	//usecases

	return &Container{
		Tables: &handlers.TableHandler{Service: tableService},
		Users: &handlers.UserHandler{Service: userService},
	}
}

func main(){
	configs.LoadEnv()
	g := gin.Default()
	g.SetTrustedProxies(nil)
	g.Use(RateLimiter())

	conn, err := supabase.Connect()
	if err != nil {
		log.Fatal("Error connecting to supabase")
	}
	defer conn.Close()

	Routes(g, *BuildContainer(conn))

	g.Run(configs.URL)
}
