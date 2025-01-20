FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY ./ ./

RUN go build -o todo-service

EXPOSE 8082

CMD ["./todo-service"]
