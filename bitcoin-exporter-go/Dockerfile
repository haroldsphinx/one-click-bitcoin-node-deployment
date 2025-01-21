FROM golang:1.23.4 

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod tidy

COPY . /app/

RUN go build -o bitcoin-exporter-go main.go

EXPOSE 8335

CMD ["./bitcoin-exporter-go"]
