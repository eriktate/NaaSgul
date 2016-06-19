package models

//PushNotificationDTO represents a new notification that should be sent to a given subscriber.
type PushNotificationDTO struct {
	NotificationDTO

	SubscriberID string `json:"subscriberId" db:"subscriber_id"`
	HasBeenSeen  bool   `json:"hasBeenSeen" db:"has_been_seen"`
}

//NewPushNotificationDTO creates a new push notification given a valid notificationDTO and a subscriber ID.
func NewPushNotificationDTO(notification NotificationDTO, subID string) *PushNotificationDTO {
	return &PushNotificationDTO{
		NotificationDTO: notification,
		SubscriberID:    subID,
		HasBeenSeen:     false,
	}
}
