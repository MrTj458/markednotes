FROM golang:1.18.1-alpine

WORKDIR /app

COPY . .

RUN go install github.com/cespare/reflex@latest
RUN go mod download
