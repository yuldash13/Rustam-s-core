package controller

import (
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/logic"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/db"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Transfer struct {
	logic logic.Transfers
}

func NewTransfer(logic logic.Transfers) *Transfer {
	return &Transfer{logic: logic}
}

func (a *App) SetTransferRoutes(r *gin.RouterGroup, c *Transfer) {
	pg := r.Group("/transfers")

	pg.GET("/:id", c.GetTransfers)
	pg.GET("/by_id/:id", c.GetTransferByID)
	pg.POST("", c.MakeTransfer)
	pg.POST("/last/:id", c.CancelTransfer)
}

func (crT *Transfer) GetTransfers(c *gin.Context) {
	idS := c.Param("id")
	id, err1 := strconv.Atoi(idS)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, err1)
		return
	}

	limitS := c.Query("limit")
	limit, err2 := strconv.Atoi(limitS)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
		return
	}

	items, err3 := crT.logic.GetTransfers(c.Request.Context(), id, limit)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't get transfer: %v", err3))
		return
	}

	var response domain.GetTransfersResponse

	for _, r := range items {
		response.Transfers = append(response.Transfers, domain.Transfer{
			ID:       r.ID,
			IDFrom:   r.IDFrom,
			IDTo:     r.IDTo,
			Currency: r.Currency,
			Value:    r.Value,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (crT *Transfer) GetTransferByID(c *gin.Context) {
	idS := c.Param("id")
	id, err1 := strconv.Atoi(idS)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, err1)
		return
	}

	item, err3 := crT.logic.GetTransferByID(c.Request.Context(), id)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't get transfer: %v", err3))
		return
	}

	var response = domain.GetTransferByID{
		Transfer: &domain.Transfer{
			ID:       item.ID,
			IDFrom:   item.IDFrom,
			IDTo:     item.IDTo,
			Currency: item.Currency,
			Value:    item.Value,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (crT *Transfer) MakeTransfer(c *gin.Context) {
	var newTransfer domain.Transfer
	err1 := c.ShouldBindJSON(&newTransfer)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("can't paste transfer json: %v", err1))
		return
	}

	id, err2 := crT.logic.MakeTransfer(c.Request.Context(), &db.Transfer{
		IDFrom:   newTransfer.IDFrom,
		IDTo:     newTransfer.IDTo,
		Currency: newTransfer.Currency,
		Value:    newTransfer.Value,
	})
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't create transfer: %v", err2))
		return
	}
	c.JSON(http.StatusCreated, id)
}

func (crT *Transfer) CancelTransfer(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err1 := crT.logic.CancelTransfer(c.Request.Context(), id)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't create transfer: %v", err1))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "transfer canceled successfully"})
}
