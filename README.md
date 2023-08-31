# Тестовое задание Avito

<!-- ToC start -->
# Содержание

1. [Описание задачи](#Описание-задачи)
1. [Реализация](#Реализация)
1. [Endpoints](#Endpoints)
1. [Запуск](#Запуск)
1. [Примеры](#Примеры)
<!-- ToC end -->

# Описание задачи

Требуется реализовать сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)
Сервис должен предоставлять HTTP API и принимать/отдавать запросы/ответы в формате JSON.
Полное описание в [TASK](TASK.md).
# Реализация

- Следование дизайну REST API.
- Подход "Чистой Архитектуры" и техника внедрения зависимости.
- Работа с фреймворком [gorilla/mux](https://github.com/gin-gonic/gin).
- Работа с СУБД Postgres с использованием ORM Gorm [Gorm](https://github.com/go-gorm/gorm).
- Конфигурация приложения - библиотека [godotenv](https://github.com/joho/godotenv).
- Файл документации к API - библиотека [Swagger](https://github.com/swaggo/swag).
- Запуск из Docker.

**Структура проекта:**
```
.
├── internal      
│   ├── server    // обработчики запросов
│   ├── service   // бизнес-логика
│   ├── database  // взаимодействие с БД
│   ├── models    // модели данных
│   ├── config    // файл конфигурации
│   └── tools     // вспомогательные функции
├── cmd           
│   └── app       // точка входа в приложение
├── schema        // SQL файлы с миграциями
├── docs          // файлы документации
├── test          // файлы тестирования
```

# Endpoints

- [POST] /history/ - получение истории получения/удаления сегментов пользователю за период времени.
    - Тело запроса:
        - user_id - уникальный идентификатор пользователя.
        - month - номер месяца.
        - year - год.

- [POST] /segment/ - создание нового сегмента.
    - Тело запроса:
        - name - название сегмента.
        - percentage - процент пользователей, которые попадут в этот сегмент автоматически.
- [DELETE] /segment/ - удаление сегмента по его названию.
    - Тело запроса:
        - name - название сегмента.
- [POST] /user/ - создание нового пользователя.
- [POST] /user/addSegment - добавляет пользователя в существующий сегмент.
    - Тело запроса:    
        - add_name - названия сегментов, которые нужно добавить.    
        - delete_name - названия сегментов, которые нужно убрать.
        - user_id - уникальный идентификатор пользователя.
        - expires_at - дата, когда нужно автоматически убрать эти сегменты.
- [GET] /user/all - получить всех пользователей.             
- [POST] /user/segment - получить сегменты, в которых состоит пользователь.   
    - Тело запроса:                                                          
        - user_id - уникальный идентификатор пользователя.                   
# Запуск

```
make build
make run
```

Если приложение запускается впервые, необходимо применить миграции к базе данных:

```
make migrate-up
```
При необходимости, можно обновить Swagger
```
make swag
```

# Примеры

Запросы сгенерированы из Postman для cURL.

### 1. [POST]  /user

**Запрос:**
```
$ curl --location --request POST 'http://localhost:8000/user'
```
**Тело ответа:**
```
{
    "user_id": 1,
}
```

### 2. [GET] /user/all_

**Запрос:**
```
$ curl --location 'http://localhost:8000/user/all'
```
**Тело ответа:**
```
[
    {
        "ID": 1,
        "Segments": null
    }
]
```

### 3. [POST] /segment_

**Запрос:**
```
$ curl --location 'http://localhost:8000/segment' \
--header 'Content-Type: application/json' \
--data '{
    "name":"AVITO_DISCOUNT"
}'
```
**Тело ответа:**
```
{
    "segment_id":1
}
```

### 4. [POST] /user/addSegment

**Запрос:**
```
$ curl --location 'http://localhost:8000/user/addSegment' \
--header 'Content-Type: application/json' \
--data '{
    "user_id":"1",
    "add_name":"AVITO_DISCOUNT"
}'
```
**Тело ответа:**
```
Success
```

### 5. [GET] /user/segment

**Запрос:**
```
$ curl --location 'http://localhost:8000/user/segment' \
--header 'Content-Type: application/json' \
--data '{
    "user_id":"1"
}'
```
**Тело ответа:**
```
["AVITO_DISCOUNT"]
```

### 6. [POST] /history

**Запрос:**
```
$ curl --location 'http://localhost:8000/history' \
--header 'Content-Type: application/json' \
--data '{
    "user_id":"1",
    "year":2023,
    "month":8
}'
```
**Тело ответа:**
```
1;AVITO_DISCOUNT;Add;2023-08-31 13:26:30.134443 +0300 MSK
```
### 7. [POST] /segment

**Запрос:**
```
$ curl --location 'http://localhost:8000/segment' \
--header 'Content-Type: application/json' \
--data '{
    "name":"AVITO_DISCOUNT",
    "percentage":50
}'
```
**Тело ответа:**
```
Success
```

