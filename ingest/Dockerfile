FROM golang:1.20-alpine as builder

WORKDIR /app
COPY go.* ./

RUN go mod download
COPY . ./

RUN go build -C cmd/server -v -o ../../build/server
CMD ["./build/server"]
