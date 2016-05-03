package entities

import (
	"time"

	"github.com/satori/go.uuid"
)

//Notification is a struct that represents data common across all types of notifications.
type Notification struct {
	notificationID   uuid.UUID
	subject          string
	body             string
	notificationType NotificationType
	createDate       time.Time
}

//NotificationType represents
type NotificationType string

//Const values for the NotificationType enum.
const (
	Text NotificationType = "text"
	HTML NotificationType = "html"
)

//NewNotification returns a brand new Notification entity.
func NewNotification(subject, body string, notificationType NotificationType) *Notification {
	notification := &Notification{}

	notification.SetSubject(subject)
	notification.SetBody(body)
	notification.SetNotificationType(notificationType)
	notification.SetCreateDate(time.Now())

	return notification
}

//BuildNotification returns a Notification entity built from existing Notification data.
func BuildNotification(notificationID uuid.UUID, subject, body string, notificationType NotificationType, createDate time.Time) *Notification {
	notification := &Notification{}

	notification.SetNotificationID(notificationID)
	notification.SetSubject(subject)
	notification.SetBody(body)
	notification.SetNotificationType(notificationType)
	notification.SetCreateDate(createDate)

	return notification
}

//NotificationID of the Notification entity.
func (n *Notification) NotificationID() uuid.UUID {
	return n.notificationID
}

//Subject of the Notification entity.
func (n *Notification) Subject() string {
	return n.subject
}

//Body of the Notification entity.
func (n *Notification) Body() string {
	return n.body
}

//NotificationType of the Notification entity.
func (n *Notification) NotificationType() NotificationType {
	return n.notificationType
}

//CreateDate of the Notification entity.
func (n *Notification) CreateDate() time.Time {
	return n.createDate
}

//SetNotificationID sets the notificationID of the Notification entity given a valid UUID.
func (n *Notification) SetNotificationID(notificationID uuid.UUID) {
	n.notificationID = notificationID
}

//BuildNotificationID attemps to set the notificationID of the Notification entity given a string representation of
//a notificationId.
func (n *Notification) BuildNotificationID(notificationID string) error {
	id := &uuid.UUID{}
	err := id.UnmarshalText([]byte(notificationID))

	if err == nil {
		n.notificationID = *id
	}
	return err
}

//SetSubject sets the subject of the Notification entity.
func (n *Notification) SetSubject(subject string) {
	n.subject = subject
}

//SetBody sets the body of the Notification entity.
func (n *Notification) SetBody(body string) {
	//TODO: May need to react differently based on NotificationType
	n.body = body
}

//SetNotificationType sets the notificationType of the NotificationEntity.
func (n *Notification) SetNotificationType(notificationType NotificationType) {
	n.notificationType = notificationType
}

//SetCreateDate sets the createDate of the Notification entity.
func (n *Notification) SetCreateDate(createDate time.Time) {
	n.createDate = createDate
}
