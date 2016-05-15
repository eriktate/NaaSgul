package handlers

import (
	"encoding/json"
	"log"
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
	router.Methods("POST").HandlerFunc(CreateNotification)
	router.Methods("GET").Path("/{id}").HandlerFunc(GetNotificationByID)
	router.Methods("GET").HandlerFunc(GetNotifications)
}

//CreateNotification handles requests to create a new Notification.
func CreateNotification(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	notification := &models.NotificationDTO{}

	err := decoder.Decode(notification)

	if err != nil {
		http.Error(w, "Failed to create notification: "+err.Error(), http.StatusBadRequest)
		return
	}

	newNotification, err := notificationService.CreateNotification(notification)

	if err != nil {
		http.Error(w, "Failed to create notification: "+err.Error(), http.StatusConflict)
		return
	}

	response, err := json.Marshal(newNotification)

	if err != nil {
		http.Error(w, "Failed to return new notification: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, response)
}

//GetNotificationByID handles requests for retrieving a specific Notification by their NotificationID.
func GetNotificationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	notification, err := notificationService.GetNotificationByID(vars["id"])

	if err != nil {
		http.Error(w, "Failed to retreive notification: "+err.Error(), http.StatusNotFound)
		return
	}

	response, err := json.Marshal(notification)
	if err != nil {
		http.Error(w, "Failed to retrieve notification: "+err.Error(), http.StatusInternalServerError)
	}

	writeJSON(w, response)
}

//GetNotifications catchs all GET requests to the /notifications resource. This handler is responsible for grabbing
//querystring parameters and building the proper request based on consumer supplied information.
func GetNotifications(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var notifications []*models.NotificationDTO
	var err error

	if subjects, ok := vars["subject"]; ok {
		notifications, err = notificationService.GetNotificationsBySubject(subjects[0])
	} else {
		log.Println("No subject supplied")
	}

	if err != nil {
		log.Println("Failed to retrive notifications: ", err)
		http.Error(w, "Failed to retrieve notification(s): "+err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(notifications)

	if err != nil {
		http.Error(w, "Failed to retrieve notification(s)"+err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, response)
}

//Temporary until I come up with a more standard way to respond with preset headers.
func writeJSON(w http.ResponseWriter, data []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Accept", "application/json")

	w.Write(data)
}
