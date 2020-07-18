FROM golang:1.14.4-stretch
WORKDIR /app
COPY . .

RUN go get -d -v ./...

CMD go run cmd/main.go
