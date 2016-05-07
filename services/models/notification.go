package models

import (
	"time"

	"github.com/eriktate/NaaSgul/entities"
)

//NotificationDTO is a data transfer object for entities.Notification
type NotificationDTO struct {
	NotificationID   string    `json:"notificationID"`
	Subject          string    `json:"subject"`
	Body             string    `json:"body"`
	NotificationType string    `json:"notificationType"`
	CreateDate       time.Time `json:"createDate"`
}

//ToEntity attempts to convert the NotificationDTO into a Notification entity.
func (dto *NotificationDTO) ToEntity() (*entities.Notification, error) {
	if len(dto.NotificationID) > 0 {
		return entities.BuildNotification(dto.NotificationID, dto.Subject, dto.Body, dto.NotificationType, dto.CreateDate)
	}

	entity := &entities.Notification{}
	entity.SetSubject(dto.Subject)
	entity.SetBody(dto.Body)
	entity.SetNotificationType(entities.NotificationType(dto.NotificationType))
	entity.SetCreateDate(dto.CreateDate)
	return entity, nil
}
