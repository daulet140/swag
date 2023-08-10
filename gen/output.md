
# <no value>
>	This is a sample server Petstore server.

## Swagger Example API

## Другие названия:
> *TODO Дополнительные "рабочие названия", "бизнес названия", "прикольные прозвища" проекта(с ОБЯЗАТЕЛЬНОЙ рафсшифровкой аббревиатур)*
- "name1"
- "name2"
- "name3"


## ВШЭП
> *TODO если есть*

ServiceID: <no value>

## Описание
> TODO

<no value>


## Ссылка на wiki

> <no value>



>TODO покрытие тестами - сменить ссылки на проект

| Ветка         | Покрытие тестами  |
| ------------- | --------------:   |
| master        | [![coverage report](https://gitlabce.1cb.kz/service/incomesv4/badges/main/coverage.svg)](https://gitlabce.1cb.kz/service/incomesv4/-/commits/master) |
| dev           | [![coverage report](https://gitlabce.1cb.kz/service/incomesv4/badges/dev/coverage.svg)](https://gitlabce.1cb.kz/service/incomesv4/-/commits/dev) |

## SWAGGER

 TEST
petstore.swagger.io/v2/swagger/index.html
 PROD
<no value>/v2/swagger/index.html

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

<no value>

## Права
> TODO если сервис не делает introduce(а должен!)
<no value>

## Данные по хранению файлов(в MINIO/на диске/...)
> TODO если хранит файлы - то как и где
<no value>
## Ограничения
>TODO
<no value>

## Куда опубликован
>TODO
Тест:
Internal: <no value>
External: <no value>
Бой:
Internal: <no value>
External: <no value>
