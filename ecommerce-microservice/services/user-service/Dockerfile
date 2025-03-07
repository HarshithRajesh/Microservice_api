FROM golang:1.24

WORKDIR /app
RUN go version
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

EXPOSE 8080 50051

CMD ["sh", "-c", "go run migrate/migrate.go && ./main"]
