# incomesv4
>TODO Официальное название проекта, если известно(с ОБЯЗАТЕЛЬНОЙ рафсшифровкой аббревиатур)

## Сервис получения доходов версии 4

# Описание
> TODO

Сервис получения доходов физ- или юрлица. Работает через vshep-payara(ВШЭП - внешний шлюз эл.правительства) для связи с Центром Распределения Трудовых Ресурсов(ЦРТР)

# Содержание

- [Другие названия](#другие-названия)
- [ВШЭП](#вшэп)
- [Ссылка на wiki](#ссылка-на-wiki)
- [Basic Concepts](#basic-concepts)
    - [Feeds](#feeds)
    - [Repositories](#repositories)
    - [Services](#services)
- [Contribution](#contribution)
- [Useful Links](#useful-links)
    - [Special Thanks](#special-thanks)
    - [Thanks to Contributors](#thanks-to-contributors)
    - [End User License Agreement (EULA)](#end-user-license-agreement-eula)

# Другие названия:
> *TODO Дополнительные "рабочие названия", "бизнес названия", "прикольные прозвища" проекта(с ОБЯЗАТЕЛЬНОЙ рафсшифровкой аббревиатур)*
- "доходы"
- "сверхновые доходы"
- "самые новые доходы"


# ВШЭП
> *TODO если есть*

ServiceID: ...

# Ссылка на wiki

> TODO

...

>TODO покрытие тестами - сменить ссылки на проект

| Ветка         | Покрытие тестами  |
| ------------- | --------------:   |
| master        | [![coverage report](https://gitlabce.1cb.kz/service/incomesv4/badges/main/coverage.svg)](https://gitlabce.1cb.kz/service/incomesv4/-/commits/master) |
| dev           | [![coverage report](https://gitlabce.1cb.kz/service/incomesv4/badges/dev/coverage.svg)](https://gitlabce.1cb.kz/service/incomesv4/-/commits/dev) |

## SWAGGER
>TODO указать свои ссылки

TEST
https://incomesv4-staging.1cb.kz/swagger/index.html
PROD
https://incomesv4.1cb.kz/swagger/index.html

Сгенерировать документацию: swag init -g cmd/main.go --parseDependency --parseInternal из корня

## Правила работы с ветками git
>TODO указать свои правила

Работу вести в ветке dev, при желании можно пользоваться flow "1 ветка - 1 задача" с последующим мерджем в dev

## Запуск проекта
>TODO указать детали

Положить актуальный conf.json в папке config проекта, рядом с default.conf.json. Допускается скопировать default.conf.json и заполнить авторизационные данные БД и auth
Для запуска - go run cmd/main.go

## Информация для сопровождения
>TODO куда сервис кладет аудит и т.д. - помощь поддержке

Сервис хранит факты запросов в таблицах individual_requests и company_requests - для физ. и юр. лиц соответственно. Статус 7 означает успешно выполненный запрос, статус 6 - ошибка запроса. Тело ответа с ВШЭП хранится в отдельной слинкованной таблице с постфиксом _data.

Используемые таблицы(текущее расположение см.в конфигах/запрашивай у DBA)

incomesv4.individual_requests - запросы на данные физ.лиц

incomesv4.individual_requests_data - тело ответа ВШЭП для запросов физ.лиц, связаны по ключу

incomesv4.individual_requests_tokens - токен КДП для запросов физ.лиц(асинхрон), связаны по ключу


incomesv4.company_requests - запросы на данные юр.лиц

incomesv4.company_requests_data - тело ответа ВШЭП для запросов юр.лиц, связаны по ключу

## Права
> TODO если сервис не делает introduce(а должен!)
Все права раздаются по схеме "incomesv2.%название типа запроса%.do"

## Данные по хранению файлов(в MINIO/на диске/...)
> TODO если хранит файлы - то как и где

## Ограничения
>TODO

Не может быть два запроса по одному ИИН/БИН и типу запроса за 10 минут.
Запрос не может быть выполнен без указания токена КДП или типа согласия(consent) - без токена мы не сможем сделать запрос, т.к это перс.данные

## Куда опубликован
>TODO

Secure2\Test2\WSO(бой\тест)?
