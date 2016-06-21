package mysql

import (
	"fmt"
	"log"

	"github.com/eriktate/NaaSgul/config"
	"github.com/eriktate/NaaSgul/entities"
	"github.com/eriktate/NaaSgul/providers/mysql/models"
	sm "github.com/eriktate/NaaSgul/services/models"
	"github.com/satori/go.uuid"
	//Need to do a blank import for sqlx
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//IMPORTANT NOTE: Service models are referenced through the "sm" alias as seen in the imports above.

//NotificationProvider is a struct that provides access to Noti
type NotificationProvider struct {
	db *sqlx.DB
}

//NewNotificationProvider returns a notification provider which fulfilles the NotificationRepo interface. If a mysql
//connection has not already been established, it will attempt to create one.
func NewNotificationProvider() (*NotificationProvider, error) {
	if db != nil {
		return &NotificationProvider{db}, nil
	}

	username := config.GetMysqlUserName()
	pword := config.GetMysqlPassword()
	hostname := config.GetMysqlHost()
	listenPort := config.GetMysqlPort()
	dbname := config.GetMysqlDatabase()

	log.Printf("Connecting to Notification Provider: %s:%s@(%s:%s)/%s", username, pword, hostname, listenPort, dbname)

	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s", username, pword, hostname, listenPort, dbname)
	err := Connect(connectionString)

	if err != nil {
		return nil, err
	}

	return &NotificationProvider{db}, nil
}

//CreateNotification adds a notification record given a Notification entity. The UUID is generated in the provider
//instead of the entity constructor to help reduce confusion about whether or not a Notification already exists in the
//database.
func (np *NotificationProvider) CreateNotification(notification *entities.Notification) (*sm.Notification, error) {
	//TODO: Need to verify which UUID generation function fits best.
	notification.SetNotificationID(uuid.NewV1())

	tx := db.MustBegin()
	tx.MustExec("call create_notification(?, ?, ?, ?, ?);", notification.NotificationID().String(), notification.Subject(), notification.Body(), string(notification.NotificationType()), notification.CreateDate())
	err := tx.Commit()

	if err != nil {
		return nil, err
	}

	return sm.BuildNotificationFromEntity(notification), nil
}

//GetNotificationByID returns the existing Notification associated with the given UUID.
func (np *NotificationProvider) GetNotificationByID(id uuid.UUID) (*sm.Notification, error) {
	notification := &models.Notification{}

	tx := db.MustBegin()
	tx.Get(notification, "call get_notification_by_id(?);", id.String())
	err := tx.Commit()

	return notification.ToServiceModel(), err
}

//GetNotificationsBySubject returns a list of Notifications that have subjects containing the given string.
func (np *NotificationProvider) GetNotificationsBySubject(subject string) []*sm.Notification {
	notifications := []*models.Notification{}

	tx := db.MustBegin()
	tx.Select(&notifications, "call get_notifications_by_subject(?);", subject)
	tx.Commit()

	serviceNotifications := make([]*sm.Notification, len(notifications))

	for _, value := range notifications {
		serviceNotifications = append(serviceNotifications, value.ToServiceModel())
	}

	return serviceNotifications
}
