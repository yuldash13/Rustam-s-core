package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	candidates "Roman_Sakutin/Roman_Sakutin/8-Service/Bank/api/swagger/accounts"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/logic"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/domain"
)

type Account struct {
	logic logic.Accounts
}

func NewAccount(logic logic.Accounts) *Account {
	return &Account{logic: logic}
}

func (a *App) SetAccountRoutes(r *gin.RouterGroup, c *Account) {
	wrapper := candidates.ServerInterfaceWrapper{
		Handler: c,
	}
	pg := r.Group("/accounts")

	pg.GET("/user/:id_user", checkUser, wrapper.GetAccounts)
	pg.GET("/:id", checkUser, wrapper.GetAccountByID)
	pg.POST("/:id_user", checkBankir, wrapper.CreateAccount)
	pg.POST("/dep/:id", checkBankir, wrapper.DepAccount)
	pg.DELETE("/:id", checkAdmin, wrapper.DeleteAccount)
}

func (crA *Account) GetAccounts(c *gin.Context, idUser candidates.AccountUserId) {
	items, err := crA.logic.GetAccounts(c.Request.Context(), idUser)
	if err != nil {
		writeUsersError(c, http.StatusInternalServerError, fmt.Errorf("can't get accounts: %v", err))
		return
	}

	response := candidates.GetAccountsResponse{
		Accounts: make([]candidates.Account, 0, len(items)),
	}

	for _, r := range items {
		response.Accounts = append(response.Accounts, candidates.Account{
			Balance:  r.Balance,
			Currency: r.Currency,
			Id:       &r.ID,
			IdUser:   &r.IDUser,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (crA *Account) GetAccountByID(c *gin.Context, id candidates.AccountId) {
	item, err := crA.logic.GetAccountByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, domain.NotFound) {
			writeUsersError(c, http.StatusNotFound, err)
			return
		}
		writeUsersError(c, http.StatusInternalServerError, fmt.Errorf("can't get account: %v", err))
		return
	}

	var response = candidates.GetAccountsByID{
		Account: candidates.Account{
			Id:       &item.ID,
			IdUser:   &item.IDUser,
			Balance:  item.Balance,
			Currency: item.Currency,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (crA *Account) CreateAccount(c *gin.Context, idUser candidates.AccountUserId) {
	var request candidates.Account
	err2 := c.ShouldBindJSON(&request)
	if err2 != nil {
		writeUsersError(c, http.StatusBadRequest, fmt.Errorf("can't paste account json: %v", err2))
		return
	}

	id, err3 := crA.logic.CreateAccount(c.Request.Context(), &domain.Account{
		IDUser:   idUser,
		Balance:  request.Balance,
		Currency: request.Currency,
	})
	if err3 != nil {
		writeUsersError(c, http.StatusInternalServerError, fmt.Errorf("can't create account: %v", err3))
		return
	}
	c.JSON(http.StatusCreated, id)
}

func (crA *Account) DepAccount(c *gin.Context, id candidates.AccountId) {
	var request candidates.DepAccount
	err2 := c.ShouldBindJSON(&request)
	if err2 != nil || request.Dep <= 0 {
		writeUsersError(c, http.StatusBadRequest, fmt.Errorf("can't paste balance json: %v", err2))
		return
	}

	err3 := crA.logic.DepAccount(c.Request.Context(), domain.DepAccount{Dep: request.Dep}, id)
	if err3 != nil {
		if errors.Is(err3, domain.NotFound) {
			writeUsersError(c, http.StatusNotFound, err3)
			return
		}
		writeUsersError(c, http.StatusInternalServerError, fmt.Errorf("can't dep account: %v", err3))
		return
	}
	c.JSON(http.StatusOK, candidates.MessageResponse{Message: "account dep successfully"})
}

func (crA *Account) DeleteAccount(c *gin.Context, id candidates.AccountId) {
	err2 := crA.logic.DeleteAccount(c.Request.Context(), id)
	if err2 != nil {
		if errors.Is(err2, domain.NotFound) {
			writeUsersError(c, http.StatusNotFound, err2)
			return
		}
		writeUsersError(c, http.StatusInternalServerError, fmt.Errorf("can't delete account: %v", err2))
		return
	}
	c.JSON(http.StatusOK, candidates.MessageResponse{Message: "account deleted successfully"})
}
