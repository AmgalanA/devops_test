from golang:1.20

WORKDIR /app

COPY main.go .

RUN go mod init example.com/app && go mod tidy

RUN go build -o app .

CMD ["./app"]