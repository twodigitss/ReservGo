package handlers

import (
	// "net/http"
	// "github.com/gin-gonic/gin"
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
