# Builder
FROM golang:alpine as builder

RUN apk --no-cache add ca-certificates git gcc musl-dev

RUN mkdir /build && mkdir /build/dist

ADD . /build/

WORKDIR /build

RUN go build -o /build/dist/bot .

# Runtime
FROM alpine:latest

RUN apk --no-cache add ca-certificates

ENV DEBUG 0
ENV TELEGRAM_TOKEN 0
ENV MONGODB_HOST localhost
ENV MONGODB_PORT 27017

COPY --from=builder /build/dist/bot /bot

CMD ["/bot"]