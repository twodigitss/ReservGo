package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/twodigitss/reserv-go/internal/modules/tables"
	"github.com/twodigitss/reserv-go/internal/shared"
)

type TableHandler struct{
	Service *tables.TableService
}
func NewTablesHandler() *TableHandler {
    return &TableHandler{}
}

func (this *TableHandler) ListAllTables(g *gin.Context){
	result, err := this.Service.ListAllTables(g.Request.Context())
	if err != nil {
		shared.JSON(g, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	shared.JSON(g, http.StatusOK, result)
}

func (this *TableHandler) FindTableById(g *gin.Context){
	var _id string = g.Param("id")
	result, err := this.Service.FindTableById(g.Request.Context(), _id)
	if err != nil {
		shared.JSON(g, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	shared.JSON(g, http.StatusOK, result)
}

func (this *TableHandler) SetTableAvailable(g *gin.Context){
	var _id string = g.Param("id")
	result, err := this.Service.SetTableAvailable(g.Request.Context(), _id)
	if err != nil {
		shared.JSON(g, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	shared.JSON(g, http.StatusOK, result)
}

func (this *TableHandler) SetTableOccupied(g *gin.Context){
	var _id string = g.Param("id")
	result, err := this.Service.SetTableOccupied(g.Request.Context(), _id)
	if err != nil {
		shared.JSON(g, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	shared.JSON(g, http.StatusOK, result)
}
