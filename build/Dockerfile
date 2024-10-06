FROM golang:1.23

COPY go.mod /app/
WORKDIR /app

RUN go mod download
COPY . .
RUN go build -o auction-systems ./cmd/main.go

EXPOSE 8080

CMD ["./auction-systems"]