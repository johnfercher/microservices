FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

EXPOSE 8080

CMD ["./userapi", "env=docker"]