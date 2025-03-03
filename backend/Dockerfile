FROM golang:1.24.0-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux go build -o graph_miner ./cmd/graph_miner.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/graph_miner .
CMD ["./graph_miner"]