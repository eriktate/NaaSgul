package main

import (
	"log"

	"github.com/eriktate/NaaSgul/providers/mysql"
	"github.com/eriktate/NaaSgul/services"
)

//NotificationProvider is the one place that has to be updated in order to switch NotificatinoProviders. If the database
//engine were to change, or the data was sources from other place, then all that has to change is the struct being
//returned here.
func NotificationProvider() services.NotificationRepo {
	provider, err := mysql.NewNotificationProvider()

	if err != nil {
		log.Println(err)
		log.Fatalln("Can not create a NotificationProvider")
	}

	return provider
}
