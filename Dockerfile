FROM golang:1.15.2-alpine3.12 as build
RUN mkdir /app
WORKDIR /app
COPY go.mod /app
COPY go.sum /app
RUN go mod download -x
COPY . /app
ENV CGO_ENABLED=0 GO111MODULE=on
RUN go build -o meditate
FROM alpine:3.11.3
COPY --from=build /app/meditate  .
ENTRYPOINT [ "./meditate" ]
