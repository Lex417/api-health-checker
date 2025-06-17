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
