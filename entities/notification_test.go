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
		notifType := "text"
		notificationID := "c48feb4f-44c1-4e83-8b1b-fb1408f0db28"
		createDate := time.Now()

		Convey("When we attempt to build a new Notification", func() {
			notification, _ := BuildNotification(notificationID, subject, body, notifType, createDate)

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

	Convey("Given a subject, body, notificationType, invalid notificationId and createDate", t, func() {
		subject := "testing"
		body := "test body"
		notifType := "text"
		notificationID := "bad id"
		createDate := time.Now()

		Convey("When we attempt to build a new Notification", func() {
			notification, err := BuildNotification(notificationID, subject, body, notifType, createDate)

			Convey("Then we should get an error back and a nil Notification", func() {
				So(notification, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestGetters(t *testing.T) {
	Convey("Given a valid Notification entity", t, func() {
		notification, _ := BuildNotification(uuid.NewV1().String(), "testing", "test body", "text", time.Now())
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

func TestSetters(t *testing.T) {
	Convey("Given a valid Notification entity", t, func() {
		notification, _ := BuildNotification(uuid.NewV1().String(), "testing", "test body", "text", time.Now())

		Convey("When attempting to set the notificationID with a UUID", func() {
			id := uuid.NewV1()
			notification.SetNotificationID(id)

			Convey("Then the notificationID should be updated", func() {
				So(notification.NotificationID(), ShouldEqual, id)
			})
		})

		Convey("When attempting to build the notificationID", func() {
			Convey("with a valid UUID string", func() {
				id := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
				err := notification.BuildNotificationID(id)

				Convey("Then the notificationID should be updated and no errors returned", func() {
					So(err, ShouldBeNil)
					So(notification.NotificationID().String(), ShouldEqual, id)
				})
			})

			Convey("with an invalid UUID string", func() {
				id := "this is bad"
				err := notification.BuildNotificationID(id)

				Convey("Then the notificationID should not be updated and an error should be returned", func() {
					So(err, ShouldNotBeNil)
					So(notification.NotificationID(), ShouldNotEqual, id)
				})
			})
		})

		Convey("When attempting to set the subject", func() {
			subject := "new subject"
			notification.SetSubject(subject)

			Convey("Then the subject should be updated", func() {
				So(notification.Subject(), ShouldEqual, subject)
			})
		})

		Convey("When attempting to set the body", func() {
			body := "new body"
			notification.SetBody(body)

			Convey("Then the body should be updated", func() {
				So(notification.Body(), ShouldEqual, body)
			})
		})

		Convey("When attempting to set the notificationType", func() {
			notifType := HTML
			notification.SetNotificationType(notifType)

			Convey("Then the notificationType should be updated", func() {
				So(notification.NotificationType(), ShouldEqual, notifType)
			})
		})

		Convey("When attempting to set the createDate", func() {
			createDate := time.Now()
			notification.SetCreateDate(createDate)

			Convey("Then the createDate should be updated", func() {
				So(notification.CreateDate().Equal(createDate), ShouldBeTrue)
			})
		})
	})
}
