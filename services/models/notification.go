package models

import (
	"time"

	"github.com/eriktate/NaaSgul/entities"
)

//NotificationDTO is a data transfer object for entities.Notification
type NotificationDTO struct {
	NotificationID   string    `json:"notificationID" db:"NotificationId"`
	Subject          string    `json:"subject" db:"Subject"`
	Body             string    `json:"body" db:"Body"`
	NotificationType string    `json:"notificationType" db:"NotificationType"`
	CreateDate       time.Time `json:"createDate" db:"CreateDate"`
}

//BuildNotificationDTOFromEntity returns the DTO representation of the given Notification entity.
func BuildNotificationDTOFromEntity(entity *entities.Notification) *NotificationDTO {
	return &NotificationDTO{
		NotificationID:   entity.NotificationID().String(),
		Subject:          entity.Subject(),
		Body:             entity.Body(),
		NotificationType: string(entity.NotificationType()),
		CreateDate:       entity.CreateDate(),
	}
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
