# HouseBook

[![Backend CI](https://github.com/e829dce/HouseBook/actions/workflows/backend.yml/badge.svg)](https://github.com/e829dce/HouseBook/actions/workflows/backend.yml)
[![Frontend CI](https://github.com/e829dce/HouseBook/actions/workflows/frontend.yml/badge.svg)](https://github.com/e829dce/HouseBook/actions/workflows/frontend.yml)
[![Storage CI](https://github.com/e829dce/HouseBook/actions/workflows/storage.yml/badge.svg)](https://github.com/e829dce/HouseBook/actions/workflows/storage.yml)
[![Scraper CI](https://github.com/e829dce/HouseBook/actions/workflows/scraper.yml/badge.svg)](https://github.com/e829dce/HouseBook/actions/workflows/scraper.yml)

Перед началом работы необходимо создать файл `.env` в корне проекта. Примерно выглядит так:
```dotenv
POSTGRES_DB=database
POSTGRES_USER=username
POSTGRES_PASSWORD=password
```
Также перед первым запуском необходимо запустить миграции.

## Запуск миграций
> [!IMPORTANT]
> Перед запуском миграций необходимо запустить контейнеры `backend` и `db`.

```shell
$ docker compose exec backend tern migrate -m /app/migrations -c /app/migrations/tern.conf
```

## Форматирование
> [!IMPORTANT]
> Перед любым коммитом необходимо подогнать весь код под единый формат.
> Проверка форматирования встроена в Dockerfile, из-за чего проект может перестать собираться.

Для бэкенда в папке `backend`:
```shell
$ go fmt ./...
```

Для документации в папке `backend`:
```shell
$ go tool swag fmt
```

Для фронтенда в папке `frontend`:
```shell
$ npm run format
```

## Сборка проекта
> [!NOTE]
> Перед сборкой бэкенда желательно сгенерировать документацию:
> ```shell
> $ go tool swag init --parseDependency --parseInternal
> ```
> Далее документация будет доступна по адресу `http://localhost/api/docs/`.

```shell
$ docker compose build
```

## Запуск проекта
```shell
$ docker compose up
```

## Остановка проекта
```shell
$ docker compose stop
```

## Просмотр проекта
После запуска проекта фронтенд должен быть доступен в браузере по адресу `http://localhost` (без каких-либо портов). По возможности будет совершён автоматический переход на HTTPS.
