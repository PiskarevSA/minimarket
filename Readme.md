## Структура проекта и назначение директорий

`oapi` - в папке находится описание эндпоинтов в формате [openapi](https://github.com/oapi-codegen/oapi-codegen). На их основе генерируется go-код при помощи инструмента [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen). Настройки генерации задаются в файле `oapi-codegen.yml`.

Документация к инструменту находится по [ссылке](https://github.com/oapi-codegen/oapi-codegen).

`sql` - в папке находятся миграции и запросы к базе данных. На их основе генерируется go-код при помощи инструмента [sqlc](https://github.com/sqlc-dev/sqlc). Настройки генерации задаются в файле `sqlc.yaml`. 

Документация к инструменту находится по [ссылке](https://docs.sqlc.dev/en/latest/).