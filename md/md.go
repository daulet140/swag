package md

import (
	"bytes"
	"encoding/json"
	"github.com/daulet140/swag"
	"log"
	"strings"
	"text/template"
)

var readMeTemplate = `
# {{if .Swagger.Info.Title}}{{.Swagger.Info.Title}}{{else}}{ServiceTitle}{{end}}
> {{if .Name}}{{.Name}} {{else}}{Service}{{end}} 

{{if .Swagger.Info.Description}}{{.Swagger.Info.Description}}{{else}}{ServiceDescription}{{end}}


# Содержание

{{if .AddName}}- [Другие названия](#другие-названия){{end}}
{{if .VshepServiceId}}- [ВШЭП](#вшэп) {{end}}
{{if .WikiPage}}- [Ссылка на wiki](#ссылка-на-wiki) {{end}}
- [Swagger](#swagger)
- [Правила работы с ветками git](#правила-работы-с-ветками-git)
{{if .Build}}- [Запуск проекта](#запуск-проекта){{end}}
- [Информация для сопровождения](#информация-для-сопровождения)
{{if .Grant}}- [Права](#права){{end}}
{{if .FileStore}}- [Данные по хранению файлов](#данные-по-хранению-файлов){{end}}
- [Ограничения](#ограничения)
- [Куда опубликован](#куда-опубликован)
- [Digital Nikolai](#digital-nikolai)

{{if .AddName}}
# Другие названия:

> *TODO Дополнительные "рабочие названия", "бизнес названия", "прикольные прозвища" проекта(с ОБЯЗАТЕЛЬНОЙ рафсшифровкой
аббревиатур)*
	{{range .AddNames}}
- {{.}}
	{{end}}
{{end}}

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



{{if .Grant}}
# Права:
> TODO если сервис не делает introduce(а должен!)
	
  {{range .Grants}}
- {{.}}
	{{end}}
{{end}}

{{if .FileStore}}
# Данные по хранению файлов

> TODO если хранит файлы - то как и где
> (в MINIO/на диске/...)

{{.FileStore}}
{{end}}
# Ограничения

> TODO

Не может быть два запроса по одному ИИН/БИН и типу запроса за 10 минут.
Запрос не может быть выполнен без указания токена КДП или типа согласия(consent) - без токена мы не сможем сделать
запрос, т.к это перс.данные

# Куда опубликован


|            | тест | бой |
|------------|------|----:|
| внутренний |  {{.Urls.InternalDev}}   | {{.Urls.InternalProd}} |
| внешний    |  {{.Urls.ExternalDev}}   | {{.Urls.ExternalProd}} |


# Digital Nikolai

> MATCH (n:Service {name:'{{.Name}}'}) RETURN n LIMIT 25
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
	jsonObj.AddNames = strings.Split(jsonObj.AddName, ", ")
	jsonObj.Grants = strings.Split(jsonObj.Grant, ", ")

	var tpl bytes.Buffer
	if err := readMeTemplate.Execute(&tpl, jsonObj); err != nil {
		log.Fatal(err)
	}

	return tpl.Bytes(), nil
}
