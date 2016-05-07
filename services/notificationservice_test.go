package services

import (
	"testing"
	"time"

	"github.com/eriktate/NaaSgul/services/mocks"
	"github.com/eriktate/NaaSgul/services/models"
	"github.com/golang/mock/gomock"
	"github.com/satori/go.uuid"
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
				mockRepo.EXPECT().CreateNotification(gomock.Any()).Return(&newDTO, nil).MinTimes(1)
				returnedDTO, err := notificationService.CreateNotification(notificationDTO)

				Convey("The NotificationRepo should be called and return no errors", func() {
					So(returnedDTO.NotificationID, ShouldEqual, newDTO.NotificationID)
					So(err, ShouldBeNil)
				})
			})

			Convey("And an invalid NotificationDTO", func() {
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

			Convey("And attempting to get notifications", func() {
				Convey("When given a valid string NotificationID", func() {
					notificationID := "c48feb4f-44c1-4e83-8b1b-fb1408f0db28"
					uuidNotificationID, _ := uuid.FromString(notificationID)
					Convey("When we call GetNotificationByID", func() {
						mockRepo.EXPECT().GetNotificationByID(uuidNotificationID).Return(&models.NotificationDTO{}, nil).MinTimes(1)
						_, err := notificationService.GetNotificationByID(notificationID)

						Convey("Then the NotificationRepo should be called and no errors returned", func() {
							So(err, ShouldBeNil)
						})
					})
				})
				Convey("When given a UUID NotificationID", func() {
					notificationID := uuid.NewV1()
					Convey("When we call GetNotificationbyID", func() {
						mockRepo.EXPECT().GetNotificationByID(notificationID).Return(&models.NotificationDTO{}, nil).MinTimes(1)
						_, err := notificationService.GetNotificationByID(notificationID)

						Convey("Then the NotificationRepo should be called and no errors returned", func() {
							So(err, ShouldBeNil)
						})
					})
				})
				Convey("When given an invalid string NotificationID", func() {
					notificationID := "bad id"
					Convey("When we call GetNotificationByID", func() {
						mockRepo.EXPECT().GetNotificationByID(gomock.Any()).MaxTimes(0)
						_, err := notificationService.GetNotificationByID(notificationID)

						Convey("Then the NotificationRepo should not be called and an error should be returned", func() {
							So(err, ShouldNotBeNil)
						})
					})
				})
				Convey("When given an invalid type as a NotificationID", func() {
					notificationID := 42
					Convey("When we call GetNotificationByID", func() {
						mockRepo.EXPECT().GetNotificationByID(gomock.Any()).MaxTimes(0)
						_, err := notificationService.GetNotificationByID(notificationID)

						Convey("Then the NotificationRepo should not be called and an error should be returned", func() {
							So(err, ShouldNotBeNil)
						})
					})
				})
				Convey("When given a subject that is at least 5 characters", func() {
					subject := "Testing"
					notifications := make([]*models.NotificationDTO, 12)

					Convey("When we call GetNotificationsBySubject", func() {
						mockRepo.EXPECT().GetNotificationsBySubject(subject).Return(notifications).MinTimes(1)
						_, err := notificationService.GetNotificationsBySubject(subject)

						Convey("Then the NotificationRepo should be called an no errors should be retuned", func() {
							So(err, ShouldBeNil)
						})
					})
				})
				Convey("When given a subject that is less than 5 characters", func() {
					subject := "Test"

					Convey("When we call GetNotificationsBySubject", func() {
						mockRepo.EXPECT().GetNotificationsBySubject(gomock.Any()).MaxTimes(0)
						_, err := notificationService.GetNotificationsBySubject(subject)

						Convey("Then the NotificationRepo should not be called and an error should be returned", func() {
							So(err, ShouldNotBeNil)
						})
					})
				})
			})
		})
	})
}
