FROM golang:alpine3.20 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/main ./cmd/app

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/bin/main ./bin/main
COPY --from=builder /app/.env .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD [ "./bin/main" ]

