# Stage 1: сборка приложения
FROM golang:1.23.1-alpine AS builder

# Устанавливаем git (если требуется для загрузки зависимостей)
RUN apk update && apk add --no-cache git

WORKDIR /app

# Копируем файлы модулей и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код проекта
COPY . .

# Собираем бинарный файл приложения.
# Поскольку main.go находится в каталоге cmd/app, указываем соответствующий путь.
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app

# Stage 2: минимальный образ для запуска
FROM alpine:latest

WORKDIR /root/

# Копируем бинарник из builder
COPY --from=builder /app/app .

# Копируем конфигурационные файлы (например, YAML-конфигурацию)
COPY --from=builder /app/configs/ ./configs/

# **Копируем файл .env в образ**
COPY .env .

EXPOSE 8080

# Запускаем приложение. Если ваше приложение само загружает .env, флаг -config не требуется.
CMD ["./app"]