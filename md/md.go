package md

import (
	"bytes"
	"encoding/json"
	"github.com/daulet140/swag"
	"log"
	"text/template"
)

var readMeTemplate = `
# {{.Name}}

## {{.FullName}}

> {{.Swagger.Info.Description}}

# Содержание

- [Другие названия](#другие-названия)
{{if .VshepServiceId}}- [ВШЭП](#вшэп) {{end}}
{{if .WikiPage}}- [Ссылка на wiki](#ссылка-на-wiki) {{end}}
- [Swagger](#swagger)
- [Правила работы с ветками git](#правила-работы-с-ветками-git)
{{if .Build}}- [Запуск проекта](#запуск-проекта){{end}}
- [Информация для сопровождения](#информация-для-сопровождения)
- [Права](#права)
- [Данные по хранению файлов](#данные-по-хранению-файлов)
- [Ограничения](#ограничения)
- [Куда опубликован](#куда-опубликован)

# Другие названия:

> *TODO Дополнительные "рабочие названия", "бизнес названия", "прикольные прозвища" проекта(с ОБЯЗАТЕЛЬНОЙ рафсшифровкой
аббревиатур)*

- "доходы"
- "сверхновые доходы"
- "самые новые доходы"
{{if .VshepServiceId}}
# ВШЭП

ServiceID: {{.VshepServiceId}}
{{end}}
{{if .WikiPage}}
# Ссылка на wiki: 
{{.WikiPage}}
{{end}}

> TODO покрытие тестами - сменить ссылки на проект

| Ветка         | Покрытие тестами  |
| ------------- | --------------:   |
| master        | [![coverage report]({{.Git}}/badges/main/coverage.svg)]({{.Git}}/-/commits/master) |
| dev           | [![coverage report]({{.Git}}/badges/dev/coverage.svg)]({{.Git}}/-/commits/dev) |

# SWAGGER

TEST
https://{{.Swagger.Host}}/swagger/index.html
PROD
https://{{.Swagger.Host}}/swagger/index.html

Сгенерировать документацию: swag init -g cmd/main.go --parseDependency --parseInternal из корня

# Правила работы с ветками git

> TODO указать свои правила

Работу вести в ветке dev, при желании можно пользоваться flow "1 ветка - 1 задача" с последующим мерджем в dev
{{if .Build}}
# Запуск проекта

> TODO указать детали

Положить актуальный conf.json в папке config проекта, рядом с default.conf.json. Допускается скопировать
default.conf.json и заполнить авторизационные данные БД и auth

Для запуска - {{.Build}}
{{end}}
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

> TODO

Secure2\Test2\WSO(бой\тест)?


|            | тест | бой |
|------------|------|----:|
| внутренний |  {{.Urls.InternalDev}}    | {{.Urls.InternalProd}}     |
| внешний    |   {{.Urls.ExternalDev}}   | {{.Urls.ExternalProd}}    |

`

func Json2MD(data []byte) ([]byte, error) {
	var jsonObj swag.Parser
	err := json.Unmarshal(data, &jsonObj)
	if err != nil {
		return nil, err
	}
	// Create a new template

	readMeTemplate, err := template.New("readme").Parse(readMeTemplate)
	if err != nil {
		log.Fatal(err)
	}

	var tpl bytes.Buffer
	if err := readMeTemplate.Execute(&tpl, jsonObj); err != nil {
		log.Fatal(err)
	}

	return tpl.Bytes(), nil
}
