package models

//PushNotification represents a new notification that should be sent to a given subscriber.
type PushNotification struct {
	Notification

	SubscriberID       string `json:"subscriberId"`
	SubscriberClientID string `json:"subscriberClientId"`
	HasBeenSeen        bool   `json:"hasBeenSeen"`
}

//NewPushNotification creates a new PushNotification model given a valid notification model and a subscriber ID.
func NewPushNotification(notification Notification, subID string) *PushNotification {
	return &PushNotification{
		Notification: notification,
		SubscriberID: subID,
		HasBeenSeen:  false,
	}
}
