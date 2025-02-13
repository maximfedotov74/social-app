APP_ENV ?= development

ENV_PATH = $(PWD)/.(APP_ENV).env

dev:
	APP_ENV=development go run cmd/main.go

build:
	go build -o bin/app cmd/main.go

prod:
	APP_ENV=production ./bin/app

make air:
	APP_ENV=development air

# -ldflags="-X 'main.Env=production'" 