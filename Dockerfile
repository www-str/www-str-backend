FROM golang:alpine3.21 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o backend ./cmd/wwwstr/

FROM alpine:3.21 AS final

# WORKDIR /app

COPY --from=builder /app/backend .

EXPOSE 8090

CMD ["./backend"]