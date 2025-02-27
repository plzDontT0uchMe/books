# Тестовое задание "мини сервис Книги"

## Описание:

```
Сервис представляет собой HTTP API сервис для работы с книгами.
``` 

### Каждая книга имеет следующие характеристики:
- Идентификатор
- Название
- Список авторов 
- Год публикации

### Каждый автор имеет следующие характеристики:
- Идентификатор
- Имя

### Функционал сервиса:
- Получение статуса сервиса (/liveness)
- Получение списка всех книг (/books)
- Получение списка всех авторов (/authors)
- Получение книги по идентификатору (/books/{id})

### Требования:
- Сервис должен быть написан на языке программирования Go
- Сервис может хранить данные в памяти
- Сервис должен быть покрыт тестами
- Сервис должен поддерживать различные уровни логирования

### Тесты:
- Написаны интеграционные тесты покрывающие взаимодействие транспортного и сервисного слоя
- Написаны модульные тесты покрывающие бизнес-логику

### Слои:
- В данном сервисе выделены 4 слоя: 
    - Транспортный
    - Сервисный
    - Репозиторий
    - Модель

### Структура проекта:
- cmd - точка входа в приложение
- app - сборка и запуск приложения
- config - конфигурация приложения
- constants - константы приложения
- dto - объекты передачи данных
- error - обработка ошибок
- interfaces - интерфейсы приложения
- models - модели данных
- rest - реализация REST API
    - author - обработка запросов авторов
    - book - обработка запросов книг
- service - бизнес-логика
- storage - хранилище данных
- berror - система ошибок
- logger - логгер
- middleware - промежуточное ПО

### Запуск приложения:
- Для быстрого запуска можно прописать команду:
- `go run ./cmd/main/.` или `make run`
- Можно указывать в env CONFIG_PATH путь к конфигурационному файлу, откуда будет браться конфигурация приложения
- Если CONFIG_PATH не будет указан, то значения берутся из переменных окружения (временно выключено, для быстрого запуска)

### Запуск тестов:
- Для запуска тестов можно прописать команду:
- `make test` или `go test ./... -cover -v` для запуска всех тестов
- `make test-service` или `go test ./internal/service -cover -v` для запуска тестов сервиса
- `make test-rest` или `go test ./internal/rest -cover -v` для запуска тестов REST API
- `make test-cover` для запуска тестов с покрытием

### Запуск с разным уровнем логирования:
- Для запуска приложения с разным уровнем логирования нужно изменить значение переменной окружения ENVIRONMENT:
- `ENVIRONMENT=dev` - минимальный уровень логирования DEBUG
- `ENVIRONMENT=stage` - минимальный уровень логирования INFO
- `ENVIRONMENT=production` - минимальный уровень логирования ERROR