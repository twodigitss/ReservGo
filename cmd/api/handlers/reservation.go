package handlers

import (
	"net/http"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/twodigitss/reserv-go/internal/modules/reservation"
	"github.com/twodigitss/reserv-go/internal/shared"
)

type ReservHandler struct{
	Service *reservation.ReservService
}
func NewReservHandler () *ReservHandler {
    return &ReservHandler{}
}

func (this *ReservHandler) Book(g *gin.Context){
	var _body reservation.HttpBodyReserv
	if err := g.ShouldBindJSON(&_body); err != nil {
		shared.JSON(g, http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var parsedBody reservation.DBReservation = reservation.DBReservation(_body)
	if parsedBody.ClientUUID == "" {
		shared.JSON(g, http.StatusBadRequest, gin.H{"message": "Invalid Client UUID "})
		return
	}
	if parsedBody.Paid == false {
		shared.JSON(g, http.StatusBadRequest, gin.H{"message": "Payment not validated"})
		return
	}

	err := this.Service.Book(g.Request.Context(), parsedBody)
	if err != nil {  
		shared.JSON(g, http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	shared.JSON(g, http.StatusOK, gin.H{"message": "Success"})
}

func (this *ReservHandler) Cancel(g *gin.Context){
	var _id string = g.Param("uuid")
	parsedID, err := uuid.Parse(_id)
	if err != nil { 
		shared.JSON(g, http.StatusBadRequest, gin.H{"message": "Invalid Reservation UUID"})
	}

	err = this.Service.Cancel(g.Request.Context(), parsedID.String())
	if err != nil {
		shared.JSON(g, http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	shared.JSON(g, http.StatusOK, gin.H{"message": "Success"})

}
