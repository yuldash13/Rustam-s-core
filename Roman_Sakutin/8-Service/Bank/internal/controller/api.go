package controller

import (
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/config"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/logic"
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	router *gin.Engine
	config *config.Config
	logic  logic.Service
	logger *zap.Logger

	user     *User
	account  *Account
	transfer *Transfer
}

func NewApp(config *config.Config, logic logic.Service, logger *zap.Logger, user *User, account *Account, transfer *Transfer) *App {
	r := gin.Default()
	return &App{
		router:   r,
		config:   config,
		logic:    logic,
		logger:   logger,
		user:     user,
		account:  account,
		transfer: transfer,
	}
}

func (a *App) StartServe() {
	a.router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "DELETE", "UPDATE"},
		AllowHeaders:     []string{"Content-Type", "User-Agent", "Content-Length"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
	}))

	api := a.router.Group("/api")
	a.SetUserRoutes(api, a.user)
	a.SetAccountRoutes(api, a.account)
	a.SetTransferRoutes(api, a.transfer)

	server := &http.Server{
		Addr:           ":" + a.config.ServerPort,
		Handler:        a.router,
		ReadTimeout:    time.Second * 30, // Get
		WriteTimeout:   time.Second * 20, //Put Post
		MaxHeaderBytes: 1 << 20,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println(err.Error())
		}
	}()
	<-ctx.Done()
}
