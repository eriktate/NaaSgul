package entities

import (
	"testing"
	"time"

	"github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreatingPushNotifications(t *testing.T) {
	Convey("Given a valid Notification entity", t, func() {
		subject := "testing"
		body := "test body"
		notifType := Text

		notification := NewNotification(subject, body, notifType)
		Convey("When NewPushNotification is called", func() {
			pushNotification := NewPushNotification(*notification)

			Convey("Then a new PushNotification entity should be returned", func() {
				So(pushNotification, ShouldNotBeNil)
				So(pushNotification.Subject(), ShouldEqual, subject)
			})
		})
	})
}

func TestPushNotificationGetters(t *testing.T) {
	Convey("Given a valid PushNotification entity", t, func() {
		pushNotification := NewPushNotification(*NewNotification("testing", "test body", Text))
		pushNotification.SetSubscriberID("b3a8ab73-41e1-4038-a27d-25f3ea64dcfc")
		pushNotification.SetSubscriberClientID("12345")
		pushNotification.SetHasBeenSeen(true)

		Convey("When getting the SubscriberID", func() {
			subscriberID := pushNotification.SubscriberID()

			Convey("The subscriberID should be returned", func() {
				So(subscriberID, ShouldNotBeNil)
			})
		})

		Convey("When getting the SubscriberClientID", func() {
			subscriberClientID := pushNotification.SubscriberClientID()

			Convey("The subscriberClientID should be returned", func() {
				So(subscriberClientID, ShouldNotBeNil)
			})
		})

		Convey("When getting the HasBeenSeen flag", func() {
			hasBeenSeen := pushNotification.HasBeenSeen()

			Convey("The hasBeenSeen flag should be returned", func() {
				So(hasBeenSeen, ShouldBeTrue)
			})
		})
	})
}

func TestPushNotificationSetters(t *testing.T) {
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
