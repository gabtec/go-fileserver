# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-fileserver .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=builder /go-fileserver /go-fileserver

ENV SHARE_FOLDER=/data
ENV PORT=8080

RUN mkdir -p /data

EXPOSE 8080

CMD ["/go-fileserver"]
