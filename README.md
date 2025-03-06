# Todo App

Простое веб-приложение на языке Go для управления списком задач (ToDo List). Реализовано с использованием фреймворка Gin и базы данных SQLite. Приложение упаковано в Docker-контейнер, а запросы проксируются с порта 80 на внутренний порт 8080 через Nginx.

---

## Описание проекта

Это учебный проект, разработанный для демонстрации навыков создания REST API на Go и развертывания в облаке. Приложение предоставляет CRUD-операции для управления задачами с валидацией данных и обработкой ошибок.

---

## Функциональность

API поддерживает следующие маршруты:

| Метод   | Маршрут         | Описание                  |
|---------|-----------------|---------------------------|
| `POST`  | `/tasks`        | Добавление новой задачи   |
| `GET`   | `/tasks`        | Получение списка задач    |
| `GET`   | `/tasks/{id}`   | Получение задачи по ID    |
| `PUT`   | `/tasks/{id}`   | Обновление задачи по ID   |
| `DELETE`| `/tasks/{id}`   | Удаление задачи по ID     |

---

### Структура задачи

Каждая задача представлена в формате JSON:

```json
{
  "id": 1,
  "title": "Buy milk",
  "description": "Go to the store",
  "completed": false
}
```

Поле title обязательно, минимум 3 символа, максимум 100.

Поле description необязательно, максимум 500 символов.

Поле completed — булево, по умолчанию false.

---

### Требования

* Go 1.22 или выше (для разработки и тестирования без Docker).
* Docker и Docker Compose (для контейнеризации и развертывания).
* SQLite (встроен через пакет github.com/mattn/go-sqlite3).

---

### Установка и запуск локально

**Через Docker (рекомендуемый способ)**
1. Склонируйте репозиторий:
```bash
git clone https://github.com/Sheritari/todo-app.git
cd todo-app
```
2. Запуск приложения:
```bash
docker-compose up --build
```

Docker соберет образ и запустит два сервиса: приложение (app) и Nginx (nginx).

API будет доступно на http://localhost:80.

3. Остановка приложения:
```bash
docker-compose down
```

**Без Docker (для разработки)**
1. Склонируйте репозиторий:
```bash
git clone https://github.com/Sheritari/todo-app.git
cd todo-app
```
2. Установка зависимостей:
```bash
go mod tidy
```
3. Запуск приложения:
```bash
go run main.go
```

API будет доступно на http://localhost:8080.

---

### Примеры запросов

1. Добавление задачи:
```bash
curl -X POST -H "Content-Type: application/json" -d "{\"title\":"\Example Title\",\"description\":\"Example Description\",\"completed\":false}" http://localhost:80/tasks
```
Ожидаемый ответ:
```json
{"id":1,"title":"Example Title","description":"Example Description","completed":false}
```
2. Список задач:
```bash
curl http://localhost:80/tasks
```
Ожидаемый ответ:
```json
[{"id":2,"title":"Example 1","description":"Example Description 1","completed":false},{"id":2,"title":"Example 2","description":"Example Description 2","completed":true}]
```
3. Задача по ID:
```bash
curl http://localhost:80/tasks/1
```
Ожидаемый ответ:
```json
{"id":1,"title":"Example Title","description":"Example Description","completed":false}
```
4. Обновление:
```bash
curl -X PUT -H "Content-Type: application/json" -d "{\"title\":\"New Example Title\",\"completed\":true}" http://localhost:80/tasks/1
```
Ожидаемый ответ:
```json
{"id":1,"title":"New Example Title","description":"","completed":true}
```
5. Удаление:
```bash
curl -X DELETE http://localhost:80/tasks/1
```
Ожидаемый ответ:
204 (No content)

---

### Обработка ошибок

* 400 Bad Request: Неверный JSON или нарушение ограничений полей.
```json
{"errors":["Field Title failed validation: min", "Field Description failed validation: max"]}
```
* 404 Not Found: Задача с указанным ID не найдена.
```json
{"error":"Task not found"}
```
* 500 Internal Server Error: Ошибка на сервере (например, проблемы с базой данных).
```json
{"error":"Failed to create task"}
```

---

### Тестирование

Проект включает unit-тесты для проверки API:
1. Перейдите в директорию handlers:
```bash
cd handlers
```
2. Запуск тестов:
```bash
go test -v
```

Тесты проверяют создание задачи, валидацию и получение списка задач.

---

# Примечания

* Проксирование: Nginx перенаправляет запросы с порта 80 на 8080 внутри Docker-сети. Порт 8080 недоступен извне.
* База данных: SQLite сохраняется в файле tasks.db, монтируемом как том.

---

### Используемые технологии

* Go: язык программирования.
* Gin: веб-фреймворк для REST API.
* SQLite: легковесная база данных.
* Docker: контейнеризация.
* Nginx: обратный прокси.
