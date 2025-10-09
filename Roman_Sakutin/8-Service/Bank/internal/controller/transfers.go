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

	pg.GET("/transfer/:id", c.GetTransferByID)
	pg.POST("", c.MakeTransfer)
	pg.POST("/last/:id", c.CancelTransfer)
}

func (crT *Transfer) GetTransferByID(c *gin.Context) {
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

	items, err3 := crT.logic.GetTransferByID(c.Request.Context(), id, limit)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't get transfer: %v", err3))
		return
	}

	var response = make([]domain.Transfer, 0, len(items))

	for _, r := range items {
		response = append(response, domain.Transfer{
			ID:        r.ID,
			IDFrom:    r.IDFrom,
			IDTo:      r.IDTo,
			TransDate: r.TransDate,
			Currency:  r.Currency,
			Value:     r.Value,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (crT *Transfer) MakeTransfer(c *gin.Context) {
	var newTransfer db.Transfer
	err1 := c.ShouldBindJSON(&newTransfer)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("can't paste transfer json: %v", err1))
		return
	}

	id, err2 := crT.logic.MakeTransfer(c.Request.Context(), &newTransfer)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't create transfer: %v", err2))
		return
	}
	c.JSON(http.StatusCreated, id)
}

func (crT *Transfer) CancelTransfer(c *gin.Context) {
	idS := c.Param("id")
	id, err1 := strconv.Atoi(idS)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, err1)
		return
	}

	t, err2 := crT.logic.CancelTransfer(c.Request.Context(), id)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't create transfer: %v", err2))
		return
	}
	c.JSON(http.StatusCreated, t.ID)
}
