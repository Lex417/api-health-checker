package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"api-health-checker/config"
	"api-health-checker/checker"
	"api-health-checker/notifier"
)

var urlStatuses = make(map[string]string)

func main() {
	configPath := flag.String("config", "config.json", "Path to the configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	interval := time.Duration(cfg.IntervalSeconds) * time.Second
	if interval <= 0 {
		interval = 60 * time.Second
		log.Printf("Warning: Invalid intervalSeconds in config. Defaulting to %v.\n", interval)
	}

	log.Printf("Starting API Health Checker. Monitoring %d URLs every %v.\n", len(cfg.URLs), interval)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	performChecks(cfg, cfg.NotificationWebhook)

	for range ticker.C {
		performChecks(cfg, cfg.NotificationWebhook)
	}
}

func performChecks(cfg *config.Config, notificationWebhook string) {
	log.Println("--- Performing health checks ---")
	for _, url := range cfg.URLs {
		statusCode, err := checker.CheckURL(url)

		previousStatus, exists := urlStatuses[url]
		currentStatus := ""
		message := ""

		if err != nil {
			currentStatus = fmt.Sprintf("DOWN (Error: %v)", err)
			message = fmt.Sprintf("❌ URL %s is DOWN! Error: %v", url, err)
		} else if statusCode >= 200 && statusCode < 300 {
			currentStatus = fmt.Sprintf("UP (Status: %d)", statusCode)
			message = fmt.Sprintf("✅ URL %s is UP (Status: %d)", url, statusCode)
		} else {
			currentStatus = fmt.Sprintf("ISSUE (Status: %d)", statusCode)
			message = fmt.Sprintf("⚠️ URL %s has an ISSUE (Status: %d)", url, statusCode)
		}

		if !exists || previousStatus != currentStatus {
			notifier.SendNotification(message, notificationWebhook)
			urlStatuses[url] = currentStatus
		} else {
			log.Printf("  %s - Current Status: %s\n", url, currentStatus)
		}
	}
	log.Println("--- Health checks complete ---")
}
