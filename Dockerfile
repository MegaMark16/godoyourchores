FROM golang:1.23.1 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -o godoyourchores
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o godoyourchores


FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/godoyourchores .
COPY templates/ ./templates
CMD ["./godoyourchores"]