# https://dev.to/plutov/docker-and-go-modules-3kkn
FROM golang:1.15.6 as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o accounts_service
EXPOSE 8080

ENTRYPOINT ["/app/accounts_service"]