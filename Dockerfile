FROM golang:1.22.2

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Make the init_db.sh script executable and run it
RUN chmod +x init_db.sh
RUN ./init_db.sh

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
