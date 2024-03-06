FROM golang:1.22.0-alpine3.18
RUN apk update
ENV TZ /usr/share/zoneinfo/Asia/Tokyo
WORKDIR /api
COPY go.mod /api
COPY go.sum /api
COPY bin/main /api
RUN go mod download
EXPOSE 8081
