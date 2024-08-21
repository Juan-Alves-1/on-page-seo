FROM golang:1.22.2

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOGOARCH=amd64 GOOS=linux go build -tags musl -ldflags '-w -extldflags "-static"' -a -installsuffix cgo -o main main.go

EXPOSE 8080

CMD ["./main"]
