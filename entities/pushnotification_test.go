package entities

import (
	"testing"

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
			pushNotification := NewPushNotification(notification)

			Convey("Then a new PushNotification entity should be returned", func() {
				So(pushNotification, ShouldNotBeNil)
				So(pushNotification.Subject(), ShouldEqual, subject)
			})
		})
	})
}

func TestPushNotificationGetters(t *testing.T) {
	Convey("Given a valid PushNotification entity", t, func() {
		pushNotification := NewPushNotification(NewNotification("testing", "test body", Text))
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
	Convey("Given a valid PushNotification entity", t, func() {
		pushNotification := NewPushNotification(NewNotification("testing", "test body", Text))
		pushNotification.SetSubscriberID("b3a8ab73-41e1-4038-a27d-25f3ea64dcfc")
		pushNotification.SetSubscriberClientID("12345")
		pushNotification.SetHasBeenSeen(true)

		Convey("When attempting to set the embedded Notification", func() {
			notification := NewNotification("NewTest", "New body", HTML)

			pushNotification.SetNotification(notification)
			Convey("Then the embedded fields should be updated", func() {
				So(pushNotification.Subject(), ShouldEqual, "NewTest")
				So(pushNotification.NotificationType(), ShouldEqual, HTML)
			})
		})

		Convey("When attempting to set the subscriberID", func() {
			id := uuid.NewV1()
			pushNotification.SetSubscriberID(id.String())

			Convey("Then the subscriberID should be updated", func() {
				So(pushNotification.SubscriberID(), ShouldEqual, id)
			})
		})

		Convey("When attempting to set the subscriberClientID", func() {
			id := "9876"
			pushNotification.SetSubscriberClientID(id)

			Convey("Then the subscriberClientID should be updated", func() {
				So(pushNotification.SubscriberClientID(), ShouldEqual, id)
			})
		})

		Convey("When attempting to set the hasBeenSeen flag", func() {
			hasBeenSeen := false
			pushNotification.SetHasBeenSeen(hasBeenSeen)

			Convey("Then the hasBeenSeen flag should be updated", func() {
				So(pushNotification.HasBeenSeen(), ShouldBeFalse)
			})
		})
	})
}
