/*
Package plivo provides message notification integration for Plivo.

Usage:

	package main

	import (
		"log"

		"github.com/casdoor/notify"
		"github.com/casdoor/notify/service/plivo"
	)

	func main() {
		plivoSvc, err := plivo.New(
			&plivo.ClientOptions{
				AuthID:    "<Your-Plivo-Auth-Id>",
				AuthToken: "<Your-Plivo-Auth-Token>",
			}, &plivo.MessageOptions{
				Source: "<Your-Plivo-Source-Number>",
			})
		if err != nil {
			log.Fatalf("plivo.New() failed: %s", err.Error())
		}

		plivoSvc.AddReceivers("Destination1")

		notifier := notify.New()
		notifier.UseServices(plivoSvc)

		err = notifier.Send(context.Background(), "subject", "message")
		if err != nil {
			log.Fatalf("notifier.Send() failed: %s", err.Error())
		}

		log.Printf("notification sent")
	}
*/
package plivo
