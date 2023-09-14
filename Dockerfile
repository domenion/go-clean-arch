# syntax=docker/dockerfile:1

# FROM golang:1.20-alpine AS build

# RUN apk update
# RUN apk add build-base 

# WORKDIR /go/delivery

# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# COPY ./src .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./out/cmd/server/app ./src/cmd/server/main.go

# RUN mkdir -p ./out/internal/configs
# COPY ./src/internal/configs/config.yaml ./out/internal/configs/

FROM alpine:latest

RUN apk update
RUN apk --no-cache add ca-certificates

ENV GO111MODULE=on
ENV TZ=Asia/Bangkok

WORKDIR /app
# COPY --from=build /go/delivery/out ./
COPY ./out ./

EXPOSE 3000

CMD ./cmd/server/app --config=./internal/configs/