FROM sinistersig/alpine-go:latest
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . /go/src/app/
