package notifier

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func SendNotification(message string, notificationWebhook string) {
	log.Printf("NOTIFICATION: %s\n", message)

	if notificationWebhook != "" {
		payload := map[string]string{"text": message}
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			log.Printf("Error marshalling webhook payload: %v\n", err)
			return
		}

		client := http.Client{Timeout: 3 * time.Second}
		resp, err := client.Post(notificationWebhook, "application/json", bytes.NewBuffer(jsonPayload))
		if err != nil {
			log.Printf("Error sending webhook notification to %s: %v\n", notificationWebhook, err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Webhook notification to %s failed with status code: %d\n", notificationWebhook, resp.StatusCode)
		} else {
			log.Printf("Webhook notification sent successfully to %s\n", notificationWebhook)
		}
	}
}
