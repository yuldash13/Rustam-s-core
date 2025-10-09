package controller

import (
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/logic"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/db"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/domain"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Account struct {
	logic logic.Accounts
}

func NewAccount(logic logic.Accounts) *Account {
	return &Account{logic: logic}
}

func (a *App) SetAccountRoutes(r *gin.RouterGroup, c *Account) {
	pg := r.Group("/accounts")

	pg.GET("/user/:id_user", c.GetAccounts)
	pg.GET("/:id", c.GetAccountByID)
	pg.POST("/:id_user", c.CreateAccount)
	pg.POST("/dep/:id", c.DepAccount)
	pg.DELETE("/:id", c.DeleteAccount)
}

func (crA *Account) GetAccounts(c *gin.Context) {
	idS := c.Param("id_user")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	items, err := crA.logic.GetAccounts(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't get accounts: %v", err))
		return
	}

	var response = make([]domain.Account, 0, len(items))

	for _, r := range items {
		response = append(response, domain.Account{
			ID:       r.ID,
			IDUser:   r.IDUser,
			Balance:  r.Balance,
			Date:     r.Date,
			Currency: r.Currency,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (crA *Account) GetAccountByID(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	item, err := crA.logic.GetAccountByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, db.NotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't get account: %v", err))
		return
	}

	var resp = domain.Account{
		ID:       item.ID,
		IDUser:   item.IDUser,
		Balance:  item.Balance,
		Date:     item.Date,
		Currency: item.Currency,
	}

	c.JSON(http.StatusOK, resp)
}

func (crA *Account) CreateAccount(c *gin.Context) {
	idS := c.Param("id_user")
	id1, err1 := strconv.Atoi(idS)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, err1)
		return
	}

	var newAccount db.Account
	err2 := c.ShouldBindJSON(&newAccount)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("can't paste account json: %v", err2))
		return
	}
	newAccount.IDUser = id1

	id, err3 := crA.logic.CreateAccount(c.Request.Context(), &newAccount)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't create account: %v", err3))
		return
	}
	c.JSON(http.StatusCreated, id)
}

func (crA *Account) DepAccount(c *gin.Context) {
	idS := c.Param("id")
	id, err1 := strconv.Atoi(idS)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, err1)
		return
	}

	var dep int
	err2 := c.ShouldBindJSON(&dep)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("can't paste balance json: %v", err2))
		return
	}

	err3 := crA.logic.DepAccount(c.Request.Context(), dep, id)
	if err3 != nil {
		if errors.Is(err3, db.NotFound) {
			c.JSON(http.StatusNotFound, err3.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't dep account: %v", err3))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "account dep successfully"})
}

func (crA *Account) DeleteAccount(c *gin.Context) {
	idS := c.Param("id")
	id, err1 := strconv.Atoi(idS)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, err1)
		return
	}

	err2 := crA.logic.DeleteAccount(c.Request.Context(), id)
	if err2 != nil {
		if errors.Is(err2, db.NotFound) {
			c.JSON(http.StatusNotFound, err2.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't delete account: %v", err2))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "account deleted successfully"})
}
