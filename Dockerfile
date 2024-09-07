FROM golang:1.23.0-alpine

WORKDIR /go/src/app 

COPY . .

EXPOSE 8080 

RUN go build -o api main.go 

CMD ["./api"]