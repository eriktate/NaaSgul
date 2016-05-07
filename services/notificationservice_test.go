package services

import (
	"testing"
	"time"

	"github.com/eriktate/NaaSgul/services/mocks"
	"github.com/eriktate/NaaSgul/services/models"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNotificationService(t *testing.T) {
	Convey("Given a NotificationService", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockRepo := mocks.NewMockNotificationRepo(mockCtrl)

		notificationService := NewNotificationService(mockrepo)

		Convey("and a valid NotificationDTO", func() {
			notificationDTO := &models.NotificationDTO{
				NotificationID:   "c48feb4f-44c1-4e83-8b1b-fb1408f0db28",
				Subject:          "Test",
				Body:             "Test",
				NotificationType: "text",
				CreateDate:       time.Now(),
			}

			Convey("When we attempt to create a notification", func() {
				entity := notificationDTO.ToEntity()
				mockRepo.EXPECT().CreateNotification(gomock.Any()).Return(entity)
				returnedEntity := notificationService.CreateNotification(notificationDTO)

				Convey("A notification entity should be created and passed to the NotificationRepo", func() {
					So(returnedEntity, ShouldEqual, entity)
				})
			})
		})
	})
}
