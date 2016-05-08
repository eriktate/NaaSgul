package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eriktate/NaaSgul/services"
	"github.com/eriktate/NaaSgul/services/models"
	"github.com/gorilla/mux"
)

var notificationService *services.NotificationService

//InitNotificationHandler takes a NotificationRepo and a SubRouter and sets up handlers and routing for working with the
//Notification resource.
func InitNotificationHandler(repo services.NotificationRepo, subRouter *mux.Router) {
	initRouter(subRouter)
	notificationService = services.NewNotificationService(repo)
}

//This is where all routes for the notification resource are defined.
func initRouter(router *mux.Router) {
	router.Methods("POST").Path("/").HandlerFunc(createNotification)
}

func createNotification(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	notification := &models.NotificationDTO{}

	err := decoder.Decode(notification)

	if err != nil {
		http.Error(w, "Failed to create notification", http.StatusConflict)
	}

	notificationService.CreateNotification(notification)
}
