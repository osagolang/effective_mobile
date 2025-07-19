# Effective Mobile Subscriptions Service

## Описание
REST API сервис для управления подписками пользователей.

---

## Быстрый старт

### 1. Клонирование репозитория
```bash
git clone <репозиторий>
cd <имя проекта>
```

### 2. Создание .env файла
Создайте файл `.env` в корне проекта с содержимым для подключения к PostgreSQL:
```
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
DB_SSLMODE=disable
```

### 3. Установка зависимостей
```bash
go mod tidy
```

### 4. Запуск базы данных (Postgres) через Docker Compose
```bash
docker-compose up -d
```

### 5. Применение миграций
```bash
go install github.com/rubenv/sql-migrate/...@latest
sql-migrate up
```

### 6. Запуск сервиса
```bash
go run ./cmd/api/main.go
```

---

## Запуск через Docker

1. Соберите образ:
```bash
docker build -t effective-mobile .
```
2. Запустите контейнер (перед этим поднять бд):
```bash
docker run --env-file .env --network host effective-mobile
```

---

## Переменные окружения
- `DB_USER` - пользователь БД
- `DB_PASSWORD` - пароль БД
- `DB_HOST` - адрес БД
- `DB_PORT` - порт БД
- `DB_NAME` - имя БД
- `DB_SSLMODE` - режим SSL

---

## Основные эндпоинты

- `POST   /api/subscriptions` - создать подписку
- `GET    /api/subscriptions/:id` - получить подписку по id
- `PUT    /api/subscriptions/:id` - обновить подписку
- `DELETE /api/subscriptions/:id` - удалить подписку
- `GET    /api/subscriptions` - получить список подписок (фильтры: user_id, service_name, price, start_date, end_date, limit, offset)
- `GET    /api/subscriptions/totalcost` - получить суммарную стоимость подписок (фильтры: start_date, end_date, service_name, user_id)

---

## Пример запроса на создание подписки
```json
POST /api/subscriptions
{
  "service_name": "YouTube Premium",
  "price": 500,
  "user_id": "60610fee-2bf1-7777-ae6f-7326e79a0cba",
  "start_date": "01-2025"
}
```

---

## Установка зависимостей вручную
```bash
go mod tidy
```

## Применение миграций вручную
```bash
go install github.com/rubenv/sql-migrate/...@latest
sql-migrate up
```