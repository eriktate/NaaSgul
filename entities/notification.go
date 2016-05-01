package entities

import "time"
//Notification is a struct that represents data common across all types of notifications.
type Notification struct {
    subject string
    body string
    notificationType NotificationType
    createDate time.Time
}

//NotificationType represents
type NotificationType int

//Const values for the NotificationType enum.
const (
    Text NotificationType = iota
    HTML
)
