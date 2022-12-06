package handlers

import "github.com/gin-gonic/gin"

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"success": true})
}
