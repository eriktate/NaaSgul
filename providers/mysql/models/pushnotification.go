package models

import "github.com/eriktate/NaaSgul/entities"

type PushNotification struct {
	Notification

	SubscriberID       string `db:"subscriber_id"`
	SubscriberClientID string `db:"subscriber_client_id"`
	HasBeenSeen        bool   `db:"has_been_seen"`
}

func BuildPushNotificationFromEntity(entity entities.PushNotification) *PushNotification {
	return &PushNotification{
		Notification:       BuildNotificationFromEntity(&entity.Notification),
		SubscriberID:       entity.SubscriberID().String(),
		SubscriberClientID: entity.SubscriberClientID(),
		HasBeenSeen:        entity.HasBeenSeen(),
	}
}
