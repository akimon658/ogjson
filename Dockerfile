FROM golang:1.18.0-bullseye AS builder

WORKDIR /go/src/ogjson
COPY . .
RUN go install

FROM gcr.io/distroless/base-debian11

COPY --from=builder /go/bin/ogjson /ogjson
CMD ["/ogjson"]
