FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o loadtester cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/loadtester .
ENTRYPOINT ["./loadtester"]
