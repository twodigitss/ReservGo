package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/twodigitss/reserv-go/internal/modules/users"
)

type UserHandler struct{
	Service *users.UserService
}

func NewUserHandler() *UserHandler {
    return &UserHandler {}
}

func (this *UserHandler) ListAllUsers(g *gin.Context){
	result, err := this.Service.ListAllUsers(g.Request.Context())

	if err != nil {
		g.IndentedJSON(http.StatusInternalServerError, gin.H{"Error":err})
	}
	 g.IndentedJSON(http.StatusOK, result)

}

func (this *UserHandler) FindUserById(c *gin.Context){
	var _id string = c.Param("uuid")

	result, err := this.Service.FindUserById(c.Request.Context(), _id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error finding": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, result)
}

func (this *UserHandler) CreateUser(c *gin.Context){
	var _body users.HttpBodyClient
	if err := c.ShouldBindJSON(&_body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// c.IndentedJSON(http.StatusOK, gin.H{"DEBUG: sent body: " : _body})
	var dto users.DBClient;
	dto.Name = _body.Name
	dto.LastName = _body.LastName
	dto.Email = _body.Email

	result, err := this.Service.CreateUser(c.Request.Context(), dto)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, result)

}
