package handlers

import (
	"fmt"
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
	var _body reservation.HttpBody
	if err := g.ShouldBindJSON(&_body); err != nil {
		shared.JSON(g, http.StatusBadRequest, nil, err)
		return
	}

	body := reservation.DBReservation(_body)
	if body.ClientUUID == "" {
		shared.JSON(g, http.StatusBadRequest, nil, fmt.Errorf("Invalid Client UUID"))
		return
	}
	if body.Paid == false {
		shared.JSON(g, http.StatusBadRequest, nil, fmt.Errorf( "Payment not validated"))
		return
	}

	_, err := this.Service.Book(g.Request.Context(), body)
	if err != nil {  
		shared.JSON(g, http.StatusInternalServerError, nil, err)
		return
	}
	shared.JSON(g, http.StatusOK, nil, nil)
}

func (this *ReservHandler) Cancel(g *gin.Context){
	var _id string = g.Param("uuid")
	parsedID, err := uuid.Parse(_id)
	if err != nil { 
		shared.JSON(g, http.StatusBadRequest, nil, err)
	}

	err = this.Service.Cancel(g.Request.Context(), parsedID.String())
	if err != nil {
		shared.JSON(g, http.StatusInternalServerError, nil, err)
		return
	}
	shared.JSON(g, http.StatusOK, nil, nil)
}

func (this *ReservHandler) Update(g *gin.Context){ 
	//updates visited field to true

	var _id string = g.Param("uuid")
	parsedID, err := uuid.Parse(_id)
	if err != nil { 
		shared.JSON(g, http.StatusBadRequest, nil, err)
	}

	err = this.Service.Update(g.Request.Context(), parsedID.String())
	if err != nil {
		shared.JSON(g, http.StatusInternalServerError, nil, err)
		return
	}
	shared.JSON(g, http.StatusOK, nil, nil)

}

func (this *ReservHandler) Exists(g *gin.Context){ 
	var _id string = g.Param("uuid")
	parsedID, err := uuid.Parse(_id)
	if err != nil { 
		shared.JSON(g, http.StatusBadRequest, nil, err)
	}

	exists, err := this.Service.Exists(g.Request.Context(), parsedID.String())
	if err != nil {
		shared.JSON(g, http.StatusInternalServerError, nil, err)
		return
	}
	shared.JSON(g, http.StatusOK, exists, nil)

}

func (this *ReservHandler) GetByID(g *gin.Context){ 
	var _id string = g.Param("uuid")
	parsedID, err := uuid.Parse(_id)
	if err != nil { 
		shared.JSON(g, http.StatusBadRequest, nil, err)
	}

	obj, err := this.Service.GetByID(g.Request.Context(), parsedID.String())
	if err != nil {
		shared.JSON(g, http.StatusInternalServerError, nil, err)
		return
	}

	response := reservation.HttpRes(obj)
	shared.JSON(g, http.StatusOK, response, nil)
}

func (this *ReservHandler) GetByClientUUID(g *gin.Context){ 
	var _id string = g.Param("uuid")
	parsedID, err := uuid.Parse(_id)
	if err != nil { 
		shared.JSON(g, http.StatusBadRequest, nil, err)
	}

	dblist, err := this.Service.GetByClientUUID(g.Request.Context(), parsedID.String())
	if err != nil {
		shared.JSON(g, http.StatusInternalServerError, nil, err)
		return
	}
	response := make([]reservation.HttpRes, len(dblist))

	for i,k := range dblist {
		response[i] = reservation.HttpRes{
			UUID: k.UUID,
			CreatedAt: k.CreatedAt,
			ClientUUID: k.ClientUUID,
			TableId: k.TableId,
			Paid: k.Paid,
			Visited: k.Visited,
			VisitedDate: k.VisitedDate,
		}
	}

	shared.JSON(g, http.StatusOK, response, nil)
}
