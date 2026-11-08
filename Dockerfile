# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod ./
# RUN go mod download (uncomment when dependencies are added)

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server

# Run stage
FROM alpine:3.21

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["/app/server"]
