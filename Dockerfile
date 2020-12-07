FROM golang:1.15 AS builder

WORKDIR /usr/src

COPY . .

RUN go build -o /usr/app/main main.go

FROM ubuntu

WORKDIR /usr/app

COPY --from=builder /usr/app .

EXPOSE 8080

CMD ["./main"]