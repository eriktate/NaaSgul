package models

import "github.com/eriktate/NaaSgul/entities"

//PushNotification represents a new notification that should be sent to a given subscriber.
type PushNotification struct {
	*Notification

	SubscriberID       string `json:"subscriberId"`
	SubscriberClientID string `json:"subscriberClientId"`
	HasBeenSeen        bool   `json:"hasBeenSeen"`
}

//NewPushNotification creates a new PushNotification model given a valid notification model and a subscriber ID.
func NewPushNotification(notification *Notification, subID string) *PushNotification {
	return &PushNotification{
		Notification: notification,
		SubscriberID: subID,
		HasBeenSeen:  false,
	}
}

//ToEntity creates a new PushNotification entity from the service model.
func (pn *PushNotification) ToEntity() (*entities.PushNotification, error) {
	notificationEntity, err := pn.Notification.ToEntity()

	if err != nil {
		return nil, err
	}

	pushEntity := entities.NewPushNotification(notificationEntity)

	pushEntity.SetSubscriberID(pn.SubscriberID)
	pushEntity.SetSubscriberClientID(pn.SubscriberClientID)
	pushEntity.SetHasBeenSeen(pn.HasBeenSeen)

	return pushEntity, nil
}
