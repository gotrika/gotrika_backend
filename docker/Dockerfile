FROM golang:1.19.5-alpine AS builder

COPY . /github.com/gotrika/gotrika_backend/
WORKDIR /github.com/gotrika/gotrika_backend/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/app ./cmd/cli/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/gotrika/gotrika_backend/.bin/app .
COPY --from=0 /github.com/gotrika/gotrika_backend/docker/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]