package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func Routes(r *gin.Engine, c Container) {
	v1 := r.Group("/v1")

	r.GET("/", func(c *gin.Context){
		c.IndentedJSON(http.StatusAccepted, gin.H{"message":"Hello Gin"})
	})

	tables := v1.Group("/tables") 
	{ tables.GET("/list", c.Tables.ListAllTables ) }

	users := v1.Group("/users")
	{ users.GET("/list", c.Users.ListAllUsers) }

}
