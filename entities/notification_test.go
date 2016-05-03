package entities

import (
	"testing"
	"time"

	"github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreatingNotifications(t *testing.T) {
	Convey("Given a subject, body, and notification type", t, func() {
		subject := "testing"
		body := "test body"
		notifType := Text

		Convey("When NewNotification is called", func() {
			notification := NewNotification(subject, body, notifType)

			Convey("Then a new Notification entity should be returned", func() {
				So(notification, ShouldNotBeNil)
				So(notification.Subject(), ShouldEqual, subject)
				So(notification.Body(), ShouldEqual, body)
				So(notification.NotificationType(), ShouldEqual, notifType)
				So(notification.CreateDate(), ShouldNotBeNil)
			})
		})
	})

	Convey("Given a subject, body, notificationType, notificationId and createDate", t, func() {
		subject := "testing"
		body := "test body"
		notifType := Text
		notificationID := uuid.NewV1()
		createDate := time.Now()

		Convey("When we attempt to build a new Notification", func() {
			notification := BuildNotification(notificationID, subject, body, notifType, createDate)

			Convey("Then an existing Notification entity should be returned", func() {
				So(notification, ShouldNotBeNil)
				So(notification.NotificationID(), ShouldNotBeNil)
				So(notification.Subject(), ShouldEqual, subject)
				So(notification.Body(), ShouldEqual, body)
				So(notification.NotificationType(), ShouldEqual, notifType)
				So(notification.CreateDate(), ShouldNotBeNil)
			})
		})
	})
}

func TestGetters(t *testing.T) {
	Convey("Given a valid Notification entity", t, func() {
		notification := BuildNotification(uuid.NewV1(), "testing", "test body", Text, time.Now())
		Convey("When getting the NotificationID", func() {
			notificationID := notification.NotificationID()

			Convey("The notificationID should be returned", func() {
				So(notificationID, ShouldNotBeNil)
			})
		})

		Convey("When getting the Subject", func() {
			subject := notification.Subject()

			Convey("The subject should be returned", func() {
				So(subject, ShouldNotBeNil)
			})
		})

		Convey("When getting the Body", func() {
			body := notification.Body()

			Convey("The body should be returned", func() {
				So(body, ShouldNotBeNil)
			})
		})

		Convey("When getting the NotificationType", func() {
			notifType := notification.NotificationType()

			Convey("The notificationType should be returned", func() {
				So(notifType, ShouldNotBeNil)
			})
		})

		Convey("When getting the CreateDate", func() {
			createDate := notification.CreateDate()

			Convey("The createDate should be returned", func() {
				So(createDate, ShouldNotBeNil)
			})
		})
	})
}
