FROM golang:1.14 as builder

WORKDIR /app

COPY . .

RUN go build -o tg-poller

FROM debian:buster as app

COPY --from=builder /app/tg-poller /usr/local/bin/tg-poller

RUN apt-get update \
  && apt-get install -y ca-certificates \
  && rm -rf /var/lib/apt/lists/*

CMD ["/usr/local/bin/tg-poller"]
