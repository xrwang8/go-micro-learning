# Build stage
FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN apk update && apk add bash gcc musl-dev
RUN go build -o main main.go
#Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

CMD ["/app/main"]