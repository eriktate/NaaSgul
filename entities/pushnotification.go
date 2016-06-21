package entities

import "github.com/satori/go.uuid"

type PushNotification struct {
	*Notification

	subscriberID       uuid.UUID
	subscriberClientID string
	hasBeenSeen        bool
}

//NewPushNotification returns a pointer to a new PushNotification entity. In order to create a PushNotification, you need at least a Notification entity.
//Depending on the type of operation that the entity will be used for, either the subscriberID or the subscriberClientID will also need set.
func NewPushNotification(notification *Notification) *PushNotification {
	return &PushNotification{Notification: notification}
}

//SetNotification sets the embedded Notification entity.
func (pn *PushNotification) SetNotification(notification *Notification) {
	pn.Notification = notification
}

//SetSubscriberID sets the subscriberID for the PushNotification. If the string given is not a valid UUID, an error will be returned.
func (pn *PushNotification) SetSubscriberID(subscriberID string) error {
	subID, err := uuid.FromString(subscriberID)
	pn.subscriberID = subID

	return err
}

//SetSubscriberClientID sets the subscriberClientID for the PushNotification. What this ID represents is determined by the end-user, but it can be used
//to send notifications to particular subscribers.
func (pn *PushNotification) SetSubscriberClientID(subscriberClientID string) {
	pn.subscriberClientID = subscriberClientID
}

//SetHasBeenSeen sets the hasBeenSeen flag for the PushNotification.
func (pn *PushNotification) SetHasBeenSeen(hasBeenSeen bool) {
	pn.hasBeenSeen = hasBeenSeen
}

//SubscriberID of the PushNotification
func (pn *PushNotification) SubscriberID() uuid.UUID {
	return pn.subscriberID
}

//SubscriberClientID of the PushNotification.
func (pn *PushNotification) SubscriberClientID() string {
	return pn.subscriberClientID
}

//HasBeenSeen flag of the PushNotification.
func (pn *PushNotification) HasBeenSeen() bool {
	return pn.hasBeenSeen
}
