# calc-service-n1nja 

## Описание проекта
**calc-service-n1nja** — это веб-сервис, который принимает арифметические выражения через HTTP-запросы и возвращает результат их вычисления. Сервис проверяет валидность выражений, обрабатывает ошибки и возвращает соответствующий HTTP-ответ.

---

## Как запустить проект
1. Убедитесь, что у вас установлен Go (версия 1.18 или выше).
2. Перейдите в директорию проекта.
3. Выполните команду:
   ```bash
   go run ./cmd/calc-service-n1nja/...
   ```
4. Сервис будет доступен по адресу http://localhost:8080.

## Примеры использования
### Успешный запрос

**Запрос:**

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}
```


**Ответ:**


```json
{
  "result": 6
}
```


### Ошибка валидации (422)
**Запрос:**

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*abc"
}
```
**Ответ:**

```json
{
  "error": "Expression is not valid"
}
```
### Внутренняя ошибка сервера (500)
**Запрос:**

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": ""
}
```
**Ответ:**

```json
{
  "error": "Internal server error"
}
```
## Структура проекта
```plaintext
calc-service-n1nja/
│
├── cmd/
│   └── calc-service-n1nja/
│       └── main.go          # Основной файл сервиса
│
├── internal/
│   └── calculator/
│       ├── calculator.go    # Логика для вычисления выражений
│       └── errors.go        # Обработка ошибок
│
├── go.mod                   # Go-модуль
└── README.md                # Документация проекта
```
## Описание API
### Endpoint: `/api/v1/calculate`
#### Метод: `POST`
### Формат запроса:

```json
{
  "expression": "выражение, которое ввёл пользователь"
}
```
### Формат ответа (успех):

```json
{
  "result": "результат выражения"
}
```
### Формат ответа (ошибка):

* Если выражение недействительно:
```json
{
  "error": "Expression is not valid"
}
```
* В случае внутренней ошибки:
```json
{
  "error": "Internal server error"
}
```
## Примечания
* Программа поддерживает базовые арифметические операции: сложение (`+`), вычитание (`-`), умножение (`*`) и деление (`/`).
* Для обработки ошибок используются HTTP-коды:
   * `200 OK` — успешный ответ.
   * `422 Unprocessable Entity` — ошибка валидации.
   * `500 Internal Server Error` — внутренняя ошибка сервера.
