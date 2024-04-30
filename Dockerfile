FROM golang:1.22.2-bookworm

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o shortener

EXPOSE 80

CMD ["./shortener"]