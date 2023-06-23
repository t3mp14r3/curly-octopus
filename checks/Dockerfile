FROM golang:1.19-alpine

WORKDIR /checks

COPY . .

RUN mkdir -p /storage

RUN go get ./...

RUN go build -o /checks/run ./cmd/main.go

CMD ["/checks/run"]
