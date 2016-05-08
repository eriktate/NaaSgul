package mysql

import (
	"fmt"

	"github.com/eriktate/NaaSgul/config"
	"github.com/eriktate/NaaSgul/entities"
	"github.com/eriktate/NaaSgul/services/models"
	"github.com/satori/go.uuid"
	//Need to do a blank import for sqlx
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

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
func (np *NotificationProvider) CreateNotification(notification *entities.Notification) (*models.NotificationDTO, error) {
	//TODO: Need to verify which UUID generation function fits best.
	notification.SetNotificationID(uuid.NewV1())

	tx := db.MustBegin()
	tx.MustExec("call create_notification(%s, %s, %s, %d, %d)", notification.NotificationID(), notification.Subject(), notification.Body(), notification.NotificationType(), notification.CreateDate())
	err := tx.Commit()

	if err != nil {
		return nil, err
	}

	return models.BuildNotificationDTOFromEntity(notification), nil
}

//TODO: Need to implement this.
func (np *NotificationProvider) GetNotificationByID(id uuid.UUID) (*models.NotificationDTO, error) {
	return &models.NotificationDTO{}, nil
}

//TODO: Need to implement this.
func (np *NotificationProvider) GetNotificationsBySubject(subject string) []*models.NotificationDTO {
	notifications := make([]*models.NotificationDTO, 2)
	return notifications
}
