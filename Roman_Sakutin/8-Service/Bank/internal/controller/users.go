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

type User struct {
	logic logic.Users
}

func NewUser(logic logic.Users) *User {
	return &User{logic: logic}
}

func (a *App) SetUserRoutes(r *gin.RouterGroup, c *User) {
	pg := r.Group("/users")

	pg.GET("", c.GetUsers)
	pg.GET("/user/:id", c.GetUserByID)
	pg.POST("", c.CreateUser)
	pg.PUT("/:id", c.UpdateUser)
}

func (crU *User) GetUsers(c *gin.Context) {
	var filters domain.UsersFilter
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("wrong parameters: %v", err))
		return
	}

	items, err := crU.logic.GetUsers(c.Request.Context(), &db.UsersFilter{
		Name:        filters.Name,
		PhoneNumber: filters.PhoneNumber,
		Mail:        filters.Mail,
		Limit:       filters.Limit,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't get users: %v", err))
		return
	}

	var response = make([]domain.User, 0, len(items))

	for _, r := range items {
		response = append(response, domain.User{
			ID:          r.ID,
			Name:        r.Name,
			PhoneNumber: r.PhoneNumber,
			Mail:        r.Mail,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (crU *User) GetUserByID(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	item, err := crU.logic.GetUserByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, db.NotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't get user: %v", err))
		return
	}

	var resp = domain.User{
		ID:          item.ID,
		Name:        item.Name,
		PhoneNumber: item.PhoneNumber,
		Mail:        item.Mail,
	}
	c.JSON(http.StatusOK, resp)
}

func (crU *User) CreateUser(c *gin.Context) {
	var newUser db.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("can't paste user json: %v", err))
		return
	}

	id, err := crU.logic.CreateUser(c.Request.Context(), &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't create user: %v", err))
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (crU *User) UpdateUser(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var newUser db.User
	err1 := c.ShouldBindJSON(&newUser)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("can't paste user json: %v", err1))
		return
	}

	err2 := crU.logic.UpdateUser(c.Request.Context(), id, &newUser)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't update user: %v", err2))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "account updated successfully"})
}
