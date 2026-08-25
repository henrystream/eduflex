# ---- Build Stage ----
FROM golang:1.26.5 AS builder

WORKDIR /app

# Cache modules
#COPY go.mod go.sum ./
COPY go.mod ./

RUN go mod download

# Copy source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o service ./cmd/notification-service
    
# ---- Run Stage ----
FROM alpine:3.19

WORKDIR /app

# Add CA certificates
#RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/service .

EXPOSE 8080

CMD ["./service"]