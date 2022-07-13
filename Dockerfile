FROM golang:1.18.4-bullseye AS builder

WORKDIR /go/src/ogjson
COPY . .

RUN CGO_ENABLED=0 go install -ldflags '-s -w'

FROM scratch

COPY --from=builder /go/bin/ogjson /ogjson
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

CMD ["/ogjson"]
