FROM node:18-alpine as client-build

WORKDIR /app

COPY ./vue .

RUN npm install

RUN npm run build

FROM golang:1.18.1-alpine as server-build

WORKDIR /app

COPY . .

RUN go install github.com/cespare/reflex@latest
RUN go mod download

RUN go build cmd/markednotes/markednotes.go

FROM alpine:latest

WORKDIR /app

COPY --from=server-build /app/markednotes .
COPY --from=client-build /app/dist ./static

CMD ["./markednotes"]
