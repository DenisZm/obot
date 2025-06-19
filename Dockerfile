# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM quay.io/projectquay/golang:1.24 AS builder
ARG TARGETARCH
WORKDIR /app
COPY . .
RUN go mod download
RUN make build TARGETARCH=$TARGETARCH

FROM scratch
WORKDIR /
COPY --from=builder /app/build/obot-linux-$TARGETARCH ./obot
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./obot", "start"]
