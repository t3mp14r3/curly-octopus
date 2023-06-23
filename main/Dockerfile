FROM golang:1.19

WORKDIR /main

COPY . .

RUN go get ./...

RUN go build -o /main/run ./cmd/main.go

CMD ["/main/run"]
