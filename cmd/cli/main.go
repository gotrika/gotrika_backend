package main

import (
	"flag"
	"fmt"

	"github.com/gotrika/gotrika_backend/internal/bootstrap"
)

func main() {
	var applicationType = flag.String(
		"app_type",
		"",
		"application type like: http, scheduler, event_worker, session_worker",
	)
	flag.Parse()
	switch appType := string(*applicationType); appType {
	case "http":
		bootstrap.RunHTTP()
	case "scheduler":
		bootstrap.RunScheduler()
	case "event_worker":
		bootstrap.RunEventWorker()
	case "session_worker":
		bootstrap.RunSessionWorker()
	default:
		fmt.Println("Invalid application type, application type must be like: http, scheduler, event_worker, session_worker")
	}
}
