
FROM golang:1.15.5-alpine3.12 as base

WORKDIR /go/src/app
COPY . .

RUN apk --no-cache add build-base=0.5-r2 git=2.26.3-r0 

RUN go build -o ./bin/semver ./

FROM alpine:3.12 as main

WORKDIR /app

RUN apk update && \
    apk --no-cache add bash curl

RUN addgroup semver && \
    adduser -D -G semver semverusr && \
    chown -R :semver /app &&\
    chmod g+sw /app

USER semverusr

COPY --chown=semverusr:semver --from=base /go/src/app/bin/semver /usr/local/bin
