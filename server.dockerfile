FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o server ./cmd/

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/server .
COPY app.env .
COPY db/migration ./db/migration

EXPOSE 8080

ENTRYPOINT [ "/app/server" ]