package routes

import (
	"first-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not get event id."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not find event."})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not register event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message" : "Event registed successfully."})
}

func cancelRegister(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not get event id."})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegister(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not register event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message" : "Event registration cancelled successfully."})
}