package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/domain"
)

const (
	authHeader   = "Authorization"
	bearerPrefix = "Bearer "
	roleKey      = "role"
)

func CheckAuthMiddleWare(c *gin.Context) {
	rawHeader := c.GetHeader(authHeader)
	if rawHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
		return
	}

	if !strings.HasPrefix(rawHeader, bearerPrefix) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
		return
	}

	token := strings.TrimSpace(rawHeader[len(bearerPrefix):])
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is empty"})
		return
	}

	c.Set(roleKey, domain.HttpUser{
		Name: "Yuldash",
		Role: domain.Role(token),
	})
	c.Next()
}

func checkUser(c *gin.Context) {
	val := c.Value(roleKey).(domain.HttpUser)
	if val.Role == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user is empty"})
		return
	}
	c.Next()
}

func checkBankir(c *gin.Context) {
	val := c.Value(roleKey).(domain.HttpUser)
	if val.Role == domain.UserRole {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "weak"})
		return
	}
	c.Next()
}

func checkAdmin(c *gin.Context) {
	val := c.Value(roleKey).(domain.HttpUser)
	if val.Role == domain.UserRole || val.Role == domain.BankirRole {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "weak"})
		return
	}
	c.Next()
}

func callbackMW(c *gin.Context) {
	c.Next()

	status := c.Writer.Status()
	if status >= http.StatusInternalServerError {
		log.Printf("request failed: status=%d path=%s err=%v", status, c.FullPath(), c.Errors.Last())
	}
}
