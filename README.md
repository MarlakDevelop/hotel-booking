# Запуск приложения

``` bash
docker-compose up --build -d
```

# Запросы

### Создание бронирования
```http request
POST http://localhost:5000/order
Content-Type: application/json

{
    "email": "admin",
    "room": "lux",
    "to": "2012-01-06",
    "from": "2012-01-04"
}
```

### Получение бронирований отфильтрованных по email
```http request
GET http://localhost:5000/orders?email=admin
```
