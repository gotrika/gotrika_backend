.PHONY:
.SILENT:
.DEFAULT_GOAL := run.http

run.http:
	go run cmd/cli/main.go -app_type=http

run.scheduler:
	go run cmd/cli/main.go -app_type=scheduler

run.session_worker:
	go run cmd/cli/main.go -app_type=session_worker

run.event_worker:
	go run cmd/cli/main.go -app_type=event_worker

compose.up:
	docker-compose -f docker-compose.yml up

swag:
	swag init -g internal/bootstrap/app.go