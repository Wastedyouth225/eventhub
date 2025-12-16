# EventHub — Go Backend Service

Мини backend-сервис для обучения Go, с REST API, middleware логирования и пагинацией.

## Особенности
- POST /events — создать событие
- GET /events — получить события с пагинацией (`?limit=10&offset=0`)
- GET /health — проверка сервера
- Middleware логирования — все запросы логируются
- In-memory storage (не требуется БД)

## Как запускать
```bash
go run ./cmd/server
```
## Как тестировать
```bash 
go test ./internal/handler
```
## Стек технологий:
Go

net/http

encoding/json

sync

time
## Цель проекта:

Практика Go backend

Понимание структуры проекта и separation of concerns

Написание REST API и middleware

Пагинация и логирование

Мини-тестирование




