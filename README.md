# Практическая работа №7
## Разработка модулей программного средства

Это реализация сервиса аутентификации. Написан на языке Go.

## Запуск
Для запуска нужно склонировать репозиторий:
```bash
git clone https://github.com/ryoeuyo/goauth
```
Далее переходим в папку с проектом
```bash
cd goauth
```
Теперь подргужаем модули и запускаем
```bash
go mod tidy && go mod download
go run ./cmd/goauth
```

Запуск был проверен только на Linux системе, на Windows запуск может отлчичатся
