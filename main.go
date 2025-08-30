package main

import (
	"net/http"

	"github.com/eduzol/event_booking_api/db"
	"github.com/eduzol/event_booking_api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.Static("/client", "./client")
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080") // localhost:8080

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	eventID, err := event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event_id": eventID})
}
