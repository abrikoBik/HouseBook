# HouseBook
Перед началом работы необходимо создать файл `.env` в корне проекта. Примерно выглядит так:
```python
POSTGRES_DB=database
POSTGRES_USER=username
POSTGRES_PASSWORD=password
```
Также перед первым запуском необходимо запустить миграции.

## Запуск миграций
> [!NOTE]
> Перед запуском миграций необходимо запустить контейнеры `backend` и `db`.

```shell
$ docker compose exec backend /go/bin/tern migrate -m /usr/src/housebook/database/migrations -c /usr/src/housebook/database/migrations/tern.conf
```

## Сборка проекта
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
