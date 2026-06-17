package main

import (
	"context"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func TriggerNotification(ctx context.Context, title, body string) {
	if ctx == nil {
		return
	}

	if runtime.IsNotificationAvailable(ctx) {
		// Use an empty ID here, but ensure the App is installed via NSIS
		// When installed, Windows uses the Registry Key to identify the app
		err := runtime.SendNotification(ctx, runtime.NotificationOptions{
			Title: title,
			Body:  body,
		})
		if err != nil {
			log.Printf("Failed to send notification: %v", err)
		}
	}
}
