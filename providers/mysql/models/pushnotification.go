package models

import "github.com/eriktate/NaaSgul/entities"

//PushNotification represents a PushNotification to be used in communicating with the MySQL database.
type PushNotification struct {
	*Notification

	SubscriberID       string `db:"subscriber_id"`
	SubscriberClientID string `db:"subscriber_client_id"`
	HasBeenSeen        bool   `db:"has_been_seen"`
}

//BuildPushNotificationFromEntity converts a PushNotification entity into a PushNotification provider model.
func BuildPushNotificationFromEntity(entity *entities.PushNotification) *PushNotification {
	return &PushNotification{
		Notification:       BuildNotificationFromEntity(entity.Notification),
		SubscriberID:       entity.SubscriberID().String(),
		SubscriberClientID: entity.SubscriberClientID(),
		HasBeenSeen:        entity.HasBeenSeen(),
	}
}
