APP_ENV ?= development

ENV_PATH = $(PWD)/.$(APP_ENV).env

include $(ENV_PATH)

DB_URL = postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable

DOWN =
UP =
FORCE =
SEQ = 

dev:
	APP_ENV=development go run cmd/main.go

build:
	go build -o bin/app cmd/main.go

.PHONY: test
test:
	make css
	make js
	@go build -o ./tmp/main cmd/main.go

css:
	npx @tailwindcss/cli -i "./web/dev/css/input.css" -o "./web/public/css/main.css" --minify

prod:
	APP_ENV=production ./bin/app

air:
	APP_ENV=development air

js:
	node esbuild.js

compose:
	docker-compose --env-file $(ENV_PATH) up -d

migrate-up:
	migrate -path ./migrations -database '$(DB_URL)' up $(UP)

migrate-down:
	migrate -path ./migrations -database '$(DB_URL)' down $(DOWN)

migrate-force:
	migrate -path ./migrations -database '$(DB_URL)' force $(FORCE)

migrate-create:
	migrate create -ext sql -dir ./migrations -seq $(SEQ)

# -ldflags="-X 'main.Env=production'" 