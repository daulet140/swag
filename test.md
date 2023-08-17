
# GetPetById
> Swagger Example API
## GetPetById income adf

> This is a sample server Petstore server.

# Содержание

- [Другие названия](#другие-названия)
- [ВШЭП](#вшэп)
- [Ссылка на wiki](#ссылка-на-wiki)
- [Swagger](#swagger)
- [Правила работы с ветками git](#правила-работы-с-ветками-git)
- [Запуск проекта](#запуск-проекта)
- [Информация для сопровождения](#информация-для-сопровождения)
- [Права](#права)
- [Данные по хранению файлов](#данные-по-хранению-файлов)
- [Ограничения](#ограничения)
- [Куда опубликован](#куда-опубликован)
- [Digital Nikolai](#digital-nikolai)


# Другие названия:

> *TODO Дополнительные "рабочие названия", "бизнес названия", "прикольные прозвища" проекта(с ОБЯЗАТЕЛЬНОЙ рафсшифровкой
аббревиатур)*

- GetPetById addname adf adf adf


# ВШЭП

ServiceID: service-pet


# Ссылка на wiki:
https://petstore.swagger.io/


> TODO покрытие тестами - сменить ссылки на проект

| Ветка         | Покрытие тестами  |
| ------------- | --------------:   |
| master        | [![coverage report](https://github.com/swaggo/swag/badges/main/coverage.svg)](https://github.com/swaggo/swag/-/commits/master) |
| dev           | [![coverage report](https://github.com/swaggo/swag/badges/dev/coverage.svg)](https://github.com/swaggo/swag/-/commits/dev) |

# SWAGGER

TEST
https://petstore.swagger.io/swagger/index.html
PROD
https://petstore.swagger.io/swagger/index.html

Сгенерировать документацию: swag init -g cmd/main.go --parseDependency --parseInternal из корня

# Правила работы с ветками git

> TODO указать свои правила

Работу вести в ветке dev, при желании можно пользоваться flow "1 ветка - 1 задача" с последующим мерджем в dev

# Запуск проекта

> TODO указать детали

Положить актуальный conf.json в папке config проекта, рядом с default.conf.json. Допускается скопировать
default.conf.json и заполнить авторизационные данные БД и auth

Для запуска - go build ./cmd/main.go

# Информация для сопровождения

> TODO куда сервис кладет аудит и т.д. - помощь поддержке

Сервис хранит факты запросов в таблицах individual_requests и company_requests - для физ. и юр. лиц соответственно.
Статус 7 означает успешно выполненный запрос, статус 6 - ошибка запроса. Тело ответа с ВШЭП хранится в отдельной
слинкованной таблице с постфиксом _data.

Используемые таблицы(текущее расположение см.в конфигах/запрашивай у DBA)

incomesv4.individual_requests - запросы на данные физ.лиц

incomesv4.individual_requests_data - тело ответа ВШЭП для запросов физ.лиц, связаны по ключу

incomesv4.individual_requests_tokens - токен КДП для запросов физ.лиц(асинхрон), связаны по ключу

incomesv4.company_requests - запросы на данные юр.лиц

incomesv4.company_requests_data - тело ответа ВШЭП для запросов юр.лиц, связаны по ключу

# Права

> TODO если сервис не делает introduce(а должен!)
> Все права раздаются по схеме "incomesv2.%название типа запроса%.do"

# Данные по хранению файлов

> TODO если хранит файлы - то как и где
> (в MINIO/на диске/...)

# Ограничения

> TODO

Не может быть два запроса по одному ИИН/БИН и типу запроса за 10 минут.
Запрос не может быть выполнен без указания токена КДП или типа согласия(consent) - без токена мы не сможем сделать
запрос, т.к это перс.данные

# Куда опубликован


|            | тест | бой |
|------------|------|----:|
| внутренний |  http://petstore.swagger-staging.io    | https://petstore.swagger.io     |
| внешний    |   http://petstore.swagger-staging.io   | https://petstore.swagger.io    |


# Digital Nikolai

> MATCH (n:Service {name:'GetPetById'}) RETURN n LIMIT 25
