package handlers

import (
	"gerenciador-2fa/internal/database"
	"github.com/gin-gonic/gin"
	"context"
	"time"
)

func HealthCheck(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Verificar conexão com o MongoDB
	err := database.GetClient().Ping(ctx, nil)
	
	if err != nil {
		c.JSON(500, gin.H{
			"status": "error",
			"message": "Database connection failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
		"timestamp": time.Now().Unix(),
	})
} 