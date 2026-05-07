FROM golang:1.21-alpine AS builder

# Устанавливаем зависимости для компиляции SQLite (CGO)
RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Собираем с поддержкой CGO
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
# Копируем статику, если она есть
COPY --from=builder /app/index.html .

EXPOSE 8080
CMD ["./main"]
