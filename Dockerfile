FROM golang:1.15.2-buster as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download -x
COPY . ./
ENV CGO_ENABLED=0 GO111MODULE=on
RUN go build -mod=readonly -v -o meditate
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/meditate /app/meditate 
CMD [ "app/meditate" ]
