# API Health Checker & Notifier

A lightweight, customizable command-line interface (CLI) application built with Go that monitors the health and availability of specified URLs or APIs. It periodically checks the status of configured endpoints and sends notifications upon detecting issues.

This project is designed to be a valuable addition to your GitHub portfolio, showcasing your proficiency in Go for building robust tooling, handling network operations, and managing configurations.

## Features

* **Configurable URLs:** Specify any number of URLs/API endpoints to monitor.

* **Customizable Interval:** Set how frequently the application checks the URLs.

* **HTTP Status Check:** Verifies the HTTP status code of each endpoint.

* **Basic Notification:** Notifies the user via the console when a URL's status changes (e.g., goes down, comes back up, or returns a non-2xx status).

* **Extensible Notification System:** Designed to be easily extended for other notification methods (e.g., webhooks, Slack, email).

* **Error Handling:** Gracefully handles network errors and invalid configurations.

## Prerequisites

Before you begin, ensure you have the following installed on your system:

* **Go (version 1.16 or higher):** You can download it from [golang.org/dl](https://golang.org/dl/).

## Installation and Usage

Follow these steps to get the API Health Checker running on your machine.

### 1. Clone the Repository

First, clone this repository to your local machine:

git clone https://github.com/your-username/api-health-checker.git
cd api-health-checker

*(Note: Replace `https://github.com/your-username/api-health-checker.git` with the actual URL of your repository once you create it.)*

### 2. Create Configuration File

The application relies on a `config.json` file to define the URLs to monitor and the checking interval. Create a file named `config.json` in the root directory of the project with the following structure:


"urls": [
"https://jsonplaceholder.typicode.com/posts/1",
"https://httpbin.org/status/200",
"https://httpbin.org/status/500",
"https://does-not-exist.example.com"
],
"intervalSeconds": 10,
"notificationWebhook": ""
}


* **`urls`**: An array of strings, where each string is a URL you want to monitor.

* **`intervalSeconds`**: The time interval (in seconds) between each health check cycle. If not provided or 0, it defaults to 60 seconds.

* **`notificationWebhook`**: (Optional) A URL to which a POST request will be sent for notifications. Currently, this field is read but not actively used for actual webhook calls in the basic implementation, but it's there for future expansion.

### 3. Run the Application

You can run the application directly using the `go run` command or build an executable.

#### Option A: Run Directly

go run main.go --config config.json

#### Option B: Build and Run Executable

go build -o api-health-checker .
./api-health-checker --config config.json

The application will start monitoring the URLs and print status updates to your console.

## Project Structure

The project is organized into logical packages for better maintainability and readability:


api-health-checker/
├── main.go               # Main application entry point
├── config/               # Package for configuration loading
│   └── config.go
├── checker/              # Package for URL health checking logic
│   └── checker.go
└── notifier/             # Package for notification logic
└── notifier.go
├── config.json           # Example configuration file
└── README.md             # This documentation file


## How It Works

1.  **Configuration Loading:** `main.go` parses command-line arguments to find the `config.json` file. The `config` package then loads and deserializes this file into a Go struct.

2.  **Scheduled Checks:** A `time.Ticker` in `main.go` triggers health checks at the specified `intervalSeconds`.

3.  **Health Checking:** For each configured URL, the `checker.CheckURL` function in the `checker` package performs an HTTP GET request. It returns the HTTP status code and any error encountered.

4.  **Status Tracking & Notification:** `main.go` maintains a map (`urlStatuses`) to track the last known status of each URL. If a URL's status changes (e.g., from `up` to `down`, or from `200 OK` to `500 Internal Server Error`), the `notifier.SendNotification` function is called to log the change to the console.

## Future Enhancements

This project provides a solid foundation. Here are some ideas for further enhancements to make it even more robust and feature-rich:

* **Advanced Notifications:**

    * Integrate with popular services like Slack, PagerDuty, Twilio (for SMS), or email (SMTP).

    * Implement the `notificationWebhook` to send actual HTTP POST requests with structured JSON payloads.

* **History & Reporting:**

    * Store historical uptime data in a local database (e.g., SQLite) or a time-series database.

    * Generate daily/weekly uptime reports.

* **Thresholds & Alerts:**

    * Set custom success status codes (e.g., only 200 is OK, or any 2xx).

    * Configure response time thresholds to trigger alerts if an API is too slow.

* **Authentication:** Add support for checking APIs that require authentication (e.g., API keys, Basic Auth, OAuth tokens).

* **Parallel Checks:** Use Go routines and wait groups to perform checks on multiple URLs concurrently, speeding up the process for many endpoints.

* **Graceful Shutdown:** Implement signal handling (e.g., `os.Interrupt`) to ensure the application shuts down cleanly.

* **More Protocols:** Extend to check other protocols like TCP ports.

* **Web Dashboard:** Build a simple web interface (using Go's `net/http` or a framework like Gin/Echo) to visualize uptime and manage configurations.

## License

This project is open-source and available under the MIT License.




