package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	candidates "Roman_Sakutin/Roman_Sakutin/8-Service/Bank/api/swagger/transfers"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/logic"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/domain"
)

type Transfer struct {
	logic logic.Transfers
}

func NewTransfer(logic logic.Transfers) *Transfer {
	return &Transfer{logic: logic}
}

func (a *App) SetTransferRoutes(r *gin.RouterGroup, c *Transfer) {
	wrapper := candidates.ServerInterfaceWrapper{
		Handler: c,
	}
	pg := r.Group("/transfers")

	pg.GET("/:id", checkUser, wrapper.GetTransfers)
	pg.GET("/by_id/:id", checkUser, wrapper.GetTransferByID)
	pg.POST("", checkBankir, wrapper.MakeTransfer)
	pg.POST("/last/:id", checkAdmin, wrapper.CancelTransfer)
}

func (crT *Transfer) GetTransfers(c *gin.Context, id candidates.TransferId, params candidates.GetTransfersParams) {
	items, err3 := crT.logic.GetTransfers(c.Request.Context(), id, int64(*params.Limit))
	if err3 != nil {
		writeUsersError(c, http.StatusInternalServerError, fmt.Errorf("can't get transfer: %v", err3))
		return
	}

	response := candidates.GetTransfersResponse{
		Transfers: make([]candidates.Transfer, 0, len(items)),
	}

	for _, r := range items {
		response.Transfers = append(response.Transfers, candidates.Transfer{
			Currency: r.Currency,
			Id:       &r.ID,
			IdFrom:   r.IDFrom,
			IdTo:     r.IDTo,
			Value:    r.Value,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (crT *Transfer) GetTransferByID(c *gin.Context, id candidates.TransferId) {
	item, err3 := crT.logic.GetTransferByID(c.Request.Context(), id)
	if err3 != nil {
		writeUsersError(c, http.StatusInternalServerError, fmt.Errorf("can't get transfer: %v", err3))
		return
	}

	var response = &candidates.GetTransferByID{
		Transfer: candidates.Transfer{
			Id:       &item.ID,
			IdFrom:   item.IDFrom,
			IdTo:     item.IDTo,
			Currency: item.Currency,
			Value:    item.Value,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (crT *Transfer) MakeTransfer(c *gin.Context) {
	var request candidates.Transfer
	err1 := c.ShouldBindJSON(&request)
	if err1 != nil {
		writeUsersError(c, http.StatusBadRequest, fmt.Errorf("can't paste transfer json: %v", err1))
		return
	}

	id, err2 := crT.logic.MakeTransfer(c.Request.Context(), &domain.Transfer{
		IDFrom:   request.IdFrom,
		IDTo:     request.IdTo,
		Currency: request.Currency,
		Value:    request.Value,
	})
	if err2 != nil {
		writeUsersError(c, http.StatusInternalServerError, fmt.Errorf("can't create transfer: %v", err2))
		return
	}
	c.JSON(http.StatusCreated, id)
}

func (crT *Transfer) CancelTransfer(c *gin.Context, id candidates.TransferId) {
	err1 := crT.logic.CancelTransfer(c.Request.Context(), id)
	if err1 != nil {
		writeUsersError(c, http.StatusInternalServerError, fmt.Errorf("can't create transfer: %v", err1))
		return
	}
	c.JSON(http.StatusOK, candidates.MessageResponse{Message: "transfer canceled successfully"})
}
