package services

import (
	"errors"
	"log"

	"github.com/eriktate/NaaSgul/entities"
	"github.com/eriktate/NaaSgul/services/models"
	"github.com/satori/go.uuid"
)

//NotificationRepo is an interface that defines what a notification repository should look like.
type NotificationRepo interface {
	CreateNotification(notification *entities.Notification) (*models.Notification, error)
	GetNotificationByID(id uuid.UUID) (*models.Notification, error)
	GetNotificationsBySubject(subject string) []*models.Notification
}

//NotificationService exposes functions for interacting with notification data.
type NotificationService struct {
	repo NotificationRepo
}

//NewNotificationService returns a pointer to a NotificationService given a NotificationRepo.
func NewNotificationService(repo NotificationRepo) *NotificationService {
	return &NotificationService{repo}
}

//CreateNotification handles all requests for creating new notifications. To help with debugging, when errors occur
//the service will also return the original Notification it was given.
func (service *NotificationService) CreateNotification(notification *models.Notification) (*models.Notification, error) {
	entity, err := notification.ToEntity()
	if err != nil {
		return notification, err
	}
	log.Println("Created notification")
	return service.repo.CreateNotification(entity)
}

//GetNotificationByID accepts either a string or UUID representation of a NotificationID. If the string is invalid or
//you pass some other type, then an error is returned.
func (service *NotificationService) GetNotificationByID(notificationID interface{}) (*models.Notification, error) {
	var id uuid.UUID
	if result, ok := notificationID.(string); ok {
		var err error
		id, err = uuid.FromString(result)

		if err != nil {
			return nil, err
		}
	} else if result, ok := notificationID.(uuid.UUID); ok {
		id = result
	} else {
		return nil, errors.New("GetNotificationByID must be passed a string or a UUID")
	}

	response, err := service.repo.GetNotificationByID(id)

	if err != nil {
		log.Printf("Failed to retrieve Notification: %s", notificationID)
		log.Printf("Error message: %s", err)
	} else {
		log.Printf("Retrieved Notification: %s", response.NotificationID)
	}

	return response, err
}

//GetNotificationsBySubject will search for all notifications that have a subject containing the given string (string
//length must be 5 characters or more).
func (service *NotificationService) GetNotificationsBySubject(subject string) ([]*models.Notification, error) {
	if len(subject) < 5 {
		return nil, errors.New("GetNOtificationsBySubject must be given a subject 5 characters or longer")
	}

	return service.repo.GetNotificationsBySubject(subject), nil
}
