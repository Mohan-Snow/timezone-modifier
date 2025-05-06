FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy the entire project
COPY . .

# Build the binary from cmd directory
RUN go build -o tzmodifier ./cmd/main.go

FROM alpine:latest

RUN apk add --no-cache tzdata dbus

COPY --from=builder /app/tzmodifier /usr/local/bin/tzmodifier

# Устанавливаем необходимые права
RUN chmod +x /usr/local/bin/tzmodifier

# Создаем volume для доступа к системным файлам
VOLUME ["/etc/localtime", "/usr/share/zoneinfo", "/etc/timezone", "/run/dbus"]

# Запускаем в интерактивном режиме с подключением stdin
ENTRYPOINT ["tzmodifier"]