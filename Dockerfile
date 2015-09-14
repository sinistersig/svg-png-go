FROM sinistersig/alpine-go:latest
RUN apk add librsvg-dev
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . /go/src/app/
