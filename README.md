# Effective Mobile Subscriptions Service

## Описание
REST API сервис для управления подписками пользователей.

---

## Быстрый старт "Из коробки"

### 1. Клонирование репозитория
```bash
git clone https://github.com/osagolang/effective_mobile.git
cd effective_mobile
```

### 2. Запустите проект одной командой
```bash
docker-compose up --build
```

### 3. Сервис будет доступен по адресу:
http://localhost:8080

### 4. Swagger-документация будет доступна по адресу:
http://localhost:8080/swagger/index.html

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

