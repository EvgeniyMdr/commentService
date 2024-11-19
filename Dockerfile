# Используем официальный образ Go для сборки
FROM golang:1.22.9-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /build

# Копируем файлы модулей и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod tidy

# Копируем все исходные файлы проекта в рабочую директорию
COPY . .

## Устанавливаем утилиту goose
#RUN go install github.com/pressly/goose/v3/cmd/goose@latest
#
## Копируем миграции в контейнер
#COPY internal/db/migrations /app/migrations


# Собираем бинарный файл
RUN go build -o commentservice ./cmd/commentservice/main.go

# Собираем бинарный файл для миграций
RUN go build -o goose-custom ./cmd/commentservice/main.go

# Используем минимальный образ для запуска
FROM alpine:3.18

# Устанавливаем рабочую директорию для финального образа
WORKDIR /app

# Копируем скомпилированный бинарник из builder
COPY --from=builder /build/commentservice .
COPY --from=builder /build/goose-custom .

# Копируем директорию с миграциями в контейнер
COPY ./internal/db/migrations /app/db/migrations

## Экспонируем порты для GRPC сервера
EXPOSE ${GRPC_PORT}

# Запускаем приложение
CMD ["./commentservice"]
