FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o build/app cmd/main.go

EXPOSE 50051

CMD ["/app/build/app"]