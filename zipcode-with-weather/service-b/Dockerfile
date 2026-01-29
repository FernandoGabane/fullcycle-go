FROM golang:1.21-alpine

RUN apk add --no-cache netcat-openbsd

WORKDIR /app

COPY . .

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

RUN go build -o app cmd/main.go

ENTRYPOINT ["./entrypoint.sh"]
