FROM golang:1.13.11-stretch as builder

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bindman-dns-file-agent main.go

FROM alpine

RUN mkdir /data \
    && apk update \
    && apk add --no-cache ca-certificates \
    && update-ca-certificates

ENV BINDMAN_DNS_MANAGER_ADDR=""
ENV BINDMAN_DNS_REVERSE_PROXY_ADDR=""
ENV BINDMAN_AGENT_CONFIG_FILE=""
ENV BINDMAN_LOG_LEVEL=""
ENV BINDMAN_PORT=""

COPY --from=builder /bindman-dns-file-agent /

CMD [ "/bindman-dns-file-agent", "agent" ]
