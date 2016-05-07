package services

import (
	"github.com/eriktate/NaaSgul/entities"
	"github.com/eriktate/NaaSgul/services/models"
	"github.com/satori/go.uuid"
)

//NotificationRepo is an interface that defines what a notification data source should look like.
type NotificationRepo interface {
	CreateNotification(notification *entities.Notification) (*models.NotificationDTO, error)
	GetNotificationByID(id uuid.UUID)
	GetNotificationsBySubject(subject string)
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
//the service will also return the original NotificationDTO it was given.
func (service *NotificationService) CreateNotification(notification *models.NotificationDTO) (*models.NotificationDTO, error) {
	entity, err := notification.ToEntity()
	if err != nil {
		return notification, err
	}

	return service.repo.CreateNotification(entity)
}
