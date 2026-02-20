package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "gin-app/internal/services"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// token, err := services.LoginService(req.Username, req.Password)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"token": token,
	// })
}

// func Profile(c *gin.Context) {
// 	userID, _ := c.Get("user_id")

// 	c.JSON(http.StatusOK, gin.H{
// 		"user_id": userID,
// 	})
// }
