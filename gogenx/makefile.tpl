lang:
	easyi18n generate --pkg=locales ../locales ../locales/locales.go

api:
	 goctl-gfapi go -api {{.serviceName}}.api -dir . --home ../template

run:
	go run {{.serviceName}}.go -f etc/{{.serviceName}}.yaml -env etc/.env

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/{{.serviceName}}_service {{.serviceName}}.go