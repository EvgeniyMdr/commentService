# Используем официальный образ Go для сборки
FROM golang:1.22.9-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /build

# Копируем файлы модулей и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod tidy

# Копируем все исходные файлы проекта в рабочую директорию
COPY . .

# Собираем бинарный файл
RUN go build -o commentservice ./cmd/commentservice/main.go

# Используем минимальный образ для запуска
FROM alpine:3.18

# Устанавливаем рабочую директорию для финального образа
WORKDIR /app

# Копируем скомпилированный бинарник из builder
COPY --from=builder /build/commentservice .

## Экспонируем порты для GRPC сервера
EXPOSE ${GRPC_PORT}

# Запускаем приложение
CMD ["./commentservice"]
