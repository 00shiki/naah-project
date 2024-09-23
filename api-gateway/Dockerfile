FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o build/app cmd/main.go

CMD ["/app/build/app"]