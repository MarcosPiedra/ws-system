FROM golang:1.24-alpine AS builder

RUN apk add build-base

WORKDIR /ws

COPY ./ ./
RUN go mod download
RUN go build -v -o api ./ws-publisher/cmd/api

FROM alpine:latest AS runtime

COPY --from=builder /ws/api /ws/api
COPY --from=builder /ws/ws-publisher/cmd/api/config.yml /ws/config.yml

CMD ["/ws/api"]