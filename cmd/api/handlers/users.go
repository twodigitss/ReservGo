package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/twodigitss/reserv-go/internal/modules/users"
	"github.com/twodigitss/reserv-go/internal/shared"
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
		shared.JSON(g, http.StatusInternalServerError, gin.H{"error":err})
		return
	}

	var res []users.HttpResClient
	for _, r := range result {
		res = append(res, users.HttpResClient{
			UUID:      r.UUID,
			CreatedAt: r.CreatedAt,
			Name:      r.Name,
			LastName:  r.LastName,
			Email:     r.Email,
		})
	}

	shared.JSON(g, http.StatusOK, res)
}

func (this *UserHandler) FindUserById(g *gin.Context){
	var _id string = g.Param("uuid")

	result, err := this.Service.FindUserById(g.Request.Context(), _id)
	if err != nil {
		shared.JSON(g, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res users.HttpResClient = users.HttpResClient{
		UUID:      result.UUID,
		CreatedAt: result.CreatedAt,
		Name:      result.Name,
		LastName:  result.LastName,
		Email:     result.Email,
	}

	shared.JSON(g, http.StatusOK, res)
}

func (this *UserHandler) CreateUser(g *gin.Context){
	var _body users.HttpBodyClient
	if err := g.ShouldBindJSON(&_body); err != nil {
		shared.JSON(g, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// shared.JSON(g, http.StatusOK, gin.H{"DEBUG: sent body: " : _body})
	var dto users.DBClient
	dto.Name = _body.Name
	dto.LastName = _body.LastName
	dto.Email = _body.Email

	result, err := this.Service.CreateUser(g.Request.Context(), dto)
	if err != nil {
		shared.JSON(g, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res users.HttpResClient = users.HttpResClient{
		UUID:      result.UUID,
		CreatedAt: result.CreatedAt,
		Name:      result.Name,
		LastName:  result.LastName,
		Email:     result.Email,
	}

	shared.JSON(g, http.StatusOK, res)

}

func (this *UserHandler) DeleteUser(g *gin.Context){
	var _id string = g.Param("uuid")
	result, err := this.Service.DeleteUser(g.Request.Context(), _id)
	if err != nil {
		shared.JSON(g, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res users.HttpResClient = users.HttpResClient{
		UUID:      result.UUID,
		CreatedAt: result.CreatedAt,
		Name:      result.Name,
		LastName:  result.LastName,
		Email:     result.Email,
	}

	shared.JSON(g, http.StatusOK, res)
}
