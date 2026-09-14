FROM golang:1.23-alpine AS builder
ENV CGO_ENABLED=0
WORKDIR /src
COPY go.mod ./
COPY apps/tui_tcp_chat ./apps/tui_tcp_chat
RUN go build -o /out/tui_tcp_chat ./apps/tui_tcp_chat

FROM alpine:3.20
RUN adduser -D -H -u 10001 appuser
COPY --from=builder /out/tui_tcp_chat /usr/local/bin/tui_tcp_chat
USER appuser
ENTRYPOINT ["tui_tcp_chat"]
