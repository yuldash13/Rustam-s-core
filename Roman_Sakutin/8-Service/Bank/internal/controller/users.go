package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	candidates "Roman_Sakutin/Roman_Sakutin/8-Service/Bank/api/swagger/users"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/logic"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/domain"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/utils"
)

type User struct {
	logic logic.Users
}

func NewUser(logic logic.Users) *User {
	return &User{logic: logic}
}

func (a *App) SetUserRoutes(r *gin.RouterGroup, c *User) {
	wrapper := candidates.ServerInterfaceWrapper{
		Handler: c,
	}
	pg := r.Group("/users")

	pg.GET("", checkBankir, wrapper.GetUsers)
	pg.GET("/user/:id", checkUser, wrapper.GetUserByID)
	pg.POST("", checkAdmin, wrapper.CreateUser)
	pg.PUT("/:id", checkBankir, wrapper.UpdateUser)
}

func (crU *User) GetUsers(c *gin.Context, params candidates.GetUsersParams) {
	items, err := crU.logic.GetUsers(c.Request.Context(), &domain.UsersFilter{
		Name:        utils.PrtTo(params.Name),
		PhoneNumber: utils.PrtTo(params.PhoneNumber),
		Mail:        utils.PrtTo(params.Mail),
		Limit:       utils.PrtTo(params.Limit),
	})
	if err != nil {
		writeUsersError(c, http.StatusInternalServerError, fmt.Errorf("can't get users: %v", err))
		return
	}

	response := candidates.GetUsersResponse{
		Users: make([]candidates.User, 0, len(items)),
	}
	for _, item := range items {
		response.Users = append(response.Users, candidates.User{
			Id:          &item.ID,
			Name:        item.Name,
			PhoneNumber: item.PhoneNumber,
			Mail:        item.Mail,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (crU *User) GetUserByID(c *gin.Context, id candidates.UserId) {
	item, items, err := crU.logic.GetUserByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, domain.NotFound) {
			writeUsersError(c, http.StatusNotFound, err)
			return
		}
		writeUsersError(c, http.StatusInternalServerError, fmt.Errorf("can't get user: %v", err))
		return
	}

	var respUser = &candidates.User{
		Id:          &item.ID,
		Name:        item.Name,
		PhoneNumber: item.PhoneNumber,
		Mail:        item.Mail,
	}

	var respAccounts candidates.GetAccountsResponse

	for _, r := range items {
		respAccounts.Accounts = append(respAccounts.Accounts, candidates.Account{
			Balance:  int64(r.Balance),
			Currency: r.Currency,
			Id:       int64(r.ID),
			IdUser:   int64(r.IDUser),
		})
	}

	var response = candidates.GetUsersByID{
		Accounts: respAccounts,
		User:     *respUser,
	}
	c.JSON(http.StatusOK, response)
}

func (crU *User) CreateUser(c *gin.Context) {
	var request candidates.CreateUserJSONRequestBody
	err := c.ShouldBindJSON(&request)
	if err != nil {
		writeUsersError(c, http.StatusBadRequest, fmt.Errorf("can't paste user json: %v", err))
		return
	}

	id, err := crU.logic.CreateUser(c.Request.Context(), &domain.User{
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Mail:        request.Mail,
	})
	if err != nil {
		writeUsersError(c, http.StatusInternalServerError, fmt.Errorf("can't create user: %v", err))
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (crU *User) UpdateUser(c *gin.Context, id candidates.UserId) {
	var request candidates.UpdateUserJSONRequestBody
	err1 := c.ShouldBindJSON(&request)
	if err1 != nil {
		writeUsersError(c, http.StatusBadRequest, fmt.Errorf("can't paste user json: %v", err1))
		return
	}

	err2 := crU.logic.UpdateUser(c.Request.Context(), &domain.User{
		ID:          id,
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Mail:        request.Mail,
	})
	if err2 != nil {
		writeUsersError(c, http.StatusInternalServerError, fmt.Errorf("can't update user: %v", err2))
		return
	}

	c.JSON(http.StatusOK, candidates.MessageResponse{Message: "account updated successfully"})
}

func writeUsersError(c *gin.Context, status int, err error) {
	c.JSON(status, candidates.Error{Message: err.Error()})
}
