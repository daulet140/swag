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
>	{{.info.description}}

## {{.info.title}}

## Другие названия:
> *TODO Дополнительные "рабочие названия", "бизнес названия", "прикольные прозвища" проекта(с ОБЯЗАТЕЛЬНОЙ рафсшифровкой аббревиатур)*
- "name1"
- "name2"
- "name3"


{{if  .info.serviceID}}
## ВШЭП
> *TODO если есть*

ServiceID: {{.info.serviceID}}
{{end}}
## Описание
> TODO

{{.info.description}}

{{if .info.wikiLink}}
## Ссылка на wiki

> {{.info.wikiLink}}

{{end}}

>TODO покрытие тестами - сменить ссылки на проект

| Ветка         | Покрытие тестами  |
| ------------- | --------------:   |
| master        | [![coverage report](https://gitlabce.1cb.kz/service/incomesv4/badges/main/coverage.svg)](https://gitlabce.1cb.kz/service/incomesv4/-/commits/master) |
| dev           | [![coverage report](https://gitlabce.1cb.kz/service/incomesv4/badges/dev/coverage.svg)](https://gitlabce.1cb.kz/service/incomesv4/-/commits/dev) |

## Ссылки на документацию
 TEST
{{.host}}{{.basePath}}/swagger/index.html
 PROD
{{.info.host_prod}}{{.basePath}}/swagger/index.html

Сгенерировать документацию: swag init -g cmd/main.go --parseDependency --parseInternal из корня

## Правила работы с ветками git
>TODO указать свои правила

Работу вести в ветке dev, при желании можно пользоваться flow "1 ветка - 1 задача" с последующим мерджем в dev

## Запуск проекта
>TODO указать детали

Положить актуальный conf.json в папке config проекта, рядом с default.conf.json. Допускается скопировать default.conf.json и заполнить авторизационные данные БД и auth
Для запуска - go run cmd/main.go

{{if .info.audit}}
## Информация для сопровождения
>TODO куда сервис кладет аудит и т.д. - помощь поддержке
{{.info.audit}}
{{end}}

{{if .info.rules}}
## Права
> TODO если сервис не делает introduce(а должен!)
{{.rules}}
{{end}}

{{if .info.s3}}
## Данные по хранению файлов(в MINIO/на диске/...)
> TODO если хранит файлы - то как и где
{{.info.s3}}
{{end}}

## Ограничения
>TODO
{{.info.limitations}}

## Куда опубликован
>TODO
Тест:
Internal: {{.internal}}
External: {{.external}}
Бой:
Internal: {{.info.internal_prod}}
External: {{.info.external_prod}}
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
