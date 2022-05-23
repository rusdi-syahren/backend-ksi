FROM golang:1.13-alpine

ENV APP_NAME=gasrem-api

RUN mkdir /app

RUN apk --no-cache add bash \
	curl \
	git \
	gcc \
	g++ \
	inotify-tools

WORKDIR /app

ARG BUILD_DEVELOPMENT=staging

COPY . .

RUN mv .env.${BUILD_DEVELOPMENT} .env

RUN go mod download

RUN go build -o main .

CMD ["./main"]