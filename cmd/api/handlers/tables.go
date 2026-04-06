package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/twodigitss/reserv-go/internal/modules/tables"
)

type TableHandler struct{
	Service *tables.TableService
}
func NewTablesHandler() *TableHandler {
    return &TableHandler{}
}

func (this *TableHandler) ListAllTables(c *gin.Context){
	result, err := this.Service.ListAllTables(c.Request.Context())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

