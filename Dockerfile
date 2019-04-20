FROM golang:1.10-alpine as builder

ARG DOCK_PKG_DIR=/go/src/github.com/crabtree/mtls-transmitter

WORKDIR ${DOCK_PKG_DIR}
COPY . ${DOCK_PKG_DIR}

RUN CGO_ENABLED=0 GOOS=linux go build -o mtls-transmitter ./cmd/transmitter

FROM alpine:3.8 as certs

RUN apk add -U --no-cache ca-certificates

FROM scratch

ARG ENV=8080

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/crabtree/mtls-transmitter/mtls-transmitter .

EXPOSE ${ENV}

ENTRYPOINT ["/mtls-transmitter"]