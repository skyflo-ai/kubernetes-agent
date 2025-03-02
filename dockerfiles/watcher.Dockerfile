# Stage 1: Build the agent
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /watcher cmd/watcher/main.go

# Stage 2: Create the final image
FROM alpine:3.18

COPY --from=builder /watcher /watcher

ENTRYPOINT ["/watcher"] 