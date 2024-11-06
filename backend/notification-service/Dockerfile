FROM golang:1.23

WORKDIR /app

COPY ./ ./

RUN go mod tidy
RUN go build -o notification-service

CMD ["./notification-service"]
