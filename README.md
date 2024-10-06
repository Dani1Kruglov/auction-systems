# Прототип системы аукциона

## Описание
Прототип системы аукциона для продажи и покупки товаров. Позволяет пользователям делать ставки, определять победителя и обрабатывать транзакции.

## Архитектура
- **DDD**: Приложение разделено на слои (домен, приложение, инфраструктура, интерфейсы). Работа с методами происходит засчет интерфейсов (интерфейсы сервисов, репозиториев)
- **gRPC API с REST Gateway**: Реализован API для взаимодействия через gRPC и REST.

## База данных
- Используется **PostgreSQL** с миграциями для управления схемой (build директория).

## Функциональность
- Пополнение баланса пользователя.
- Создание и закрытие аукционов.
- Обработка ставок и транзакций.
- Уведомления участников.

## Docker
- Приложение упаковано с использованием Docker и Docker Compose.
- Запуск
    1.	Склонируйте репозиторий.
    2.	Перейдите в корневую директорию проекта.
    3.	Запустите приложение с помощью Docker: 
    #### `cd build && docker-compose up -d`

## Примеры запросов

### Регистрация пользователя
Регистрация нового пользователя:
```bash
curl --location 'http://localhost:8080/v1/users/register' \
--header 'Content-Type: application/json' \
--data '{
"name": "Dan",
"email": "khgdf",
"password": "df",
"role": "buyer",
"notification_url": "https://webhook.site/6443c3ba-4121-40b5-87e7-b38d1092e246"
}'
```

### Создание лота
Создание нового лота:
```bash
curl --location 'http://localhost:8080/v1/lots/create' \
--header 'Content-Type: application/json' \
--data '{
    "seller_id": 1,
    "item_name": "Test Item",
    "description": "oasjfush",
    "starting_price": 100.0
}'

```

### Создание аукциона
Создание нового аукциона:
```bash
curl --location 'http://localhost:8080/v1/auctions/create' \
--header 'Content-Type: application/json' \
--data '{
    "lot_id": 1,
    "start_time": "2024-10-06T21:00:00Z",
    "end_time": "2024-10-06T21:03:00Z",
    "min_step": 10.0
}'
```

### Подача ставки
Подача ставки на аукцион:
```bash
curl --location 'http://localhost:8080/v1/auctions/4/bids' \
--header 'Content-Type: application/json' \
--data '{
    "buyer_id": 10,
    "bid_amount": 100
}'
```

### Пополнение баланса пользователя
Пополнение баланса пользователя:
```bash
curl --location 'http://localhost:8080/v1/users/11/balance' \
--header 'Content-Type: application/json' \
--data '{
    "amount": 10000
}'
```