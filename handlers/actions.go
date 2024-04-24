package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"stat4market.com/db"
	"stat4market.com/models"
)

func CreateEvent(ctx *gin.Context) {
	event := models.Event{}
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.AddEvent(event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
