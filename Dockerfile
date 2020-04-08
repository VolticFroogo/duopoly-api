FROM golang:1.14.1

WORKDIR /go/src/github.com/VolticFroogo/duopoly-api
COPY . .
RUN go build -o main .

CMD ["./main"]
