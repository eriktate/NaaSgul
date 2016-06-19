package models

import (
	"time"

	"github.com/eriktate/NaaSgul/entities"
	"github.com/eriktate/NaaSgul/services/models"
)

//Notification is a provider model for entities.Notification
type Notification struct {
	NotificationID   string `json:"notificationID" db:"NotificationId"`
	Subject          string `json:"subject" db:"Subject"`
	Body             string `json:"body" db:"Body"`
	NotificationType string `json:"notificationType" db:"NotificationType"`
	CreateDate       int64  `json:"createDate" db:"CreateDate"`
}

//BuildNotificationFromEntity returns the provider model representation of the given Notification entity.
func BuildNotificationFromEntity(entity *entities.Notification) *Notification {
	return &Notification{
		NotificationID:   entity.NotificationID().String(),
		Subject:          entity.Subject(),
		Body:             entity.Body(),
		NotificationType: string(entity.NotificationType()),
		CreateDate:       entity.CreateDate().Unix(),
	}
}

//ToServiceModel converts NotificationDTO provider model
func (n *Notification) ToServiceModel(models.Notification) {
	return &models.Notification{
		NotificationID:   n.NotificationID,
		NotificationType: n.NotificationType,
		Subject:          n.Subject,
		Body:             n.Body,
		CreateDate:       time.Unix(CreateDate, 0),
	}
}
