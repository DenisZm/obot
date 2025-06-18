# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM quay.io/projectquay/golang:1.24 AS builder
ARG TARGETARCH
WORKDIR /app
COPY . .
RUN go mod download
RUN make build TARGETARCH=$TARGETARCH

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/build/obot-linux-$TARGETARCH ./obot
RUN chmod +x obot
ENTRYPOINT ["/app/obot", "start"]
