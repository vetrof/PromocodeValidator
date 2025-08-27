
# PromocodeValidator

Учебный проект, демонстрирующий, как строить сервисы на Go с гибридной архитектурой: **Vertical Slice Architecture** + **Clean Architecture**.  

Цель — показать, как организовать код так, чтобы каждая фича (например, применение промокода) была изолированной, легко тестируемой и расширяемой, а архитектура оставалась прозрачной.

---

## 🏗 Архитектурные принципы

### Vertical Slice Architecture
Каждая фича инкапсулирует в себе всё необходимое: хендлер, use case, DTO, работу с доменной моделью и доступ к данным.  
Это убирает «слоёную лапшу» и позволяет думать **фичами**, а не слоями.

### Clean Architecture
Зависимости направлены внутрь:
- `domain` ничего не знает о БД, HTTP и инфраструктуре;  
- `use case` работает через интерфейсы, а не через конкретные реализации;  
- адаптеры (репозитории, контроллеры) подключаются снаружи.  

Такой подход упрощает тестирование и развитие проекта.

---

## 📂 Структура проекта

```text
internal/
  domain/             # бизнес-сущности (PromoCode, ValidationResult)
  app/
    apply_code/       # vertical slice: применение промокода
      handler.go      # HTTP-хендлер
      usecase.go      # бизнес-логика (Application Service)
      input.go        # DTO запроса
      output.go       # DTO ответа
  adapters/
    fake/             # in-memory / фейковая реализация репозитория
    postgres/         # заготовка под PostgreSQL
pkg/
  logger/             # простой логгер
cmd/
  server/             # точка входа HTTP-сервера
````

---

## 🚀 Запуск

```bash
go run ./cmd/server
```

Сервис поднимется на `http://localhost:8080`.

---

## 📌 Пример запроса

### Применение промокода

```http
POST http://localhost:8080/promocode/apply/
Content-Type: application/json

{
  "code": "bob",
  "name": "nik",
  "phone": "777"
}
```

### Пример ответа

```json
{
  "code": "bob",
  "exists": true,
  "onTime": true,
  "applied": true,
  "appliedNow": true
}
```

Повторный вызов для того же промокода вернёт `appliedNow: false`.

---

## 🧪 Тестирование

```bash
go test ./...
```

---

## 🎯 Цель проекта

* Показать, как совмещать **вертикальные срезы** и **чистую архитектуру** на Go.
* Дать базовый шаблон для старта новых сервисов.
* Облегчить обучение структурированию кода и построению feature-based приложений.
* На практике показать, как изолировать бизнес-логику, отделить доменные модели от инфраструктуры и сделать код удобным для поддержки.

---

## 📖 Полезные ссылки

* [Vertical Slice Architecture — Jimmy Bogard](https://jimmybogard.com/vertical-slice-architecture/)
* [Clean Architecture — Uncle Bob](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)

```

