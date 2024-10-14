from golang:1.23.2

WORKDIR /go/src/goapi

COPY . .

EXPOSE 8080

RUN go build -o main cmd/main.go

CMD ["./main"]
