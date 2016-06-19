package models

import (
	"time"

	"github.com/eriktate/NaaSgul/entities"
)

//Notification is a data transfer object for entities.Notification
type Notification struct {
	NotificationID   string    `json:"notificationID" db:"NotificationId"`
	Subject          string    `json:"subject" db:"Subject"`
	Body             string    `json:"body" db:"Body"`
	NotificationType string    `json:"notificationType" db:"NotificationType"`
	CreateDate       time.Time `json:"createDate" db:"CreateDate"`
}

//BuildNotificationFromEntity returns the DTO representation of the given Notification entity.
func BuildNotificationFromEntity(entity *entities.Notification) *Notification {
	return &NotificationDTO{
		NotificationID:   entity.NotificationID().String(),
		Subject:          entity.Subject(),
		Body:             entity.Body(),
		NotificationType: string(entity.NotificationType()),
		CreateDate:       entity.CreateDate(),
	}
}

//ToEntity attempts to convert the NotificationDTO into a Notification entity.
func (dto *Notification) ToEntity() (*entities.Notification, error) {
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
