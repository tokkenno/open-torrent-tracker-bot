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

ENV TELEGRAM_TOKEN 0

COPY --from=builder /build/dist/bot /bot

CMD ["/bot"]