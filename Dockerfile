FROM sinistersig/alpine-go:latest
RUN mkdir -p /go/src/app
RUN apk add librsvg-dev
WORKDIR /go/src/app

COPY . /go/src/app/
