FROM golang:1.16-alpine3.13
WORKDIR /app
COPY . .
RUN apk update && apk add bash gcc musl-dev
RUN go build -o main main.go

CMD ["/app/main"]