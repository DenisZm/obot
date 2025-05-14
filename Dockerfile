# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM quay.io/projectquay/golang:1.24 AS builder
ARG TARGETOS
ARG TARGETARCH
ARG VERSION
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o obot -ldflags "-s -w -X=github.com/deniszm/obot/cmd.appVersion=${VERSION}"

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/obot ./obot
RUN chmod +x obot
ENTRYPOINT ["/app/obot", "start"]
