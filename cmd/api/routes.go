package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// General route handling
func Routes(r *gin.Engine, c Container) {
	v1 := r.Group("/v1")

	tables := v1.Group("/tables")
	{	tables.GET("/list", c.Tables.ListAllTables )
		tables.GET("/find/:id", c.Tables.FindTableById)
		tables.PUT("/set-available/:id", c.Tables.SetTableAvailable)
		tables.PUT("/set-occupied/:id", c.Tables.SetTableOccupied)
	}

	users := v1.Group("/users")
	{	users.GET("/list", c.Users.ListAllUsers)
		users.GET("/find/:uuid", c.Users.FindUserById)
		users.POST("/create", c.Users.CreateUser)
		users.DELETE("/delete/:uuid", c.Users.DeleteUser)
	}

	reserv := v1.Group("/reservation")
	{	reserv.POST("/book", c.Reservation.Book)
		reserv.PATCH("/update/:uuid", c.Reservation.Update)
		reserv.GET("/cancel/:uuid", c.Reservation.Cancel)
		reserv.GET("/exist/:uuid", c.Reservation.Exists)
		reserv.GET("/get-by-uuid/:uuid", c.Reservation.GetByID)
		reserv.GET("/get-by-client-uuid/:uuid", c.Reservation.GetByClientUUID)

	}

	//TODO: make it html template, idk
	r.GET("/", func(c *gin.Context) {
		routes := r.Routes() // []gin.RouteInfo

		var endpoints []gin.H
		for _, route := range routes {
			endpoints = append(
				endpoints,
				gin.H{
					"method": route.Method,
					"path": route.Path,
				},
			)
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"1ReservGo": "Made with Gin!",
			"Endpoints": endpoints,
		})

	})

}
