FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o main ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 5004
ENTRYPOINT ["./main"]
