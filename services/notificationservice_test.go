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

		notificationService := NewNotificationService(mockRepo)

		Convey("And a valid NotificationDTO", func() {
			notificationDTO := &models.NotificationDTO{
				Subject:          "Test",
				Body:             "Test",
				NotificationType: "text",
				CreateDate:       time.Now(),
			}

			newDTO := *notificationDTO
			newDTO.NotificationID = "c48feb4f-44c1-4e83-8b1b-fb1408f0db28"

			Convey("When we attempt to create a notification", func() {
				mockRepo.EXPECT().CreateNotification(gomock.Any()).Return(&newDTO, nil)
				returnedDTO, err := notificationService.CreateNotification(notificationDTO)

				Convey("The NotificationRepo should be called and returned no errors", func() {
					So(returnedDTO.NotificationID, ShouldEqual, newDTO.NotificationID)
					So(err, ShouldBeNil)
				})
			})

			Convey("and an invalid NotificationDTO", func() {
				notificationDTO := &models.NotificationDTO{
					NotificationID:   "bad ID",
					Subject:          "Test",
					Body:             "Test",
					NotificationType: "text",
					CreateDate:       time.Now(),
				}

				Convey("When we attempt to create a notification", func() {
					mockRepo.EXPECT().CreateNotification(gomock.Any()).Return(notificationDTO, nil).MaxTimes(0)
					returnedDTO, err := notificationService.CreateNotification(notificationDTO)

					Convey("Then the NotificationRepo should not be called, the original DTO should be returned, and an error should be returned", func() {
						So(returnedDTO, ShouldEqual, notificationDTO)
						So(err, ShouldNotBeNil)
					})
				})
			})
		})
	})
}
