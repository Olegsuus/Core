include .env

export ENV
export CONFIG_PATH
export POSTGRES_HOST
export POSTGRES_PORT
export POSTGRES_DB
export POSTGRES_USER
export POSTGRES_PASSWORD

.PHONY: build run docker-build docker-run migrate

build:
	@echo "Сборка приложения..."
	go build -o  .cmd/app/

run:
	@echo "Запуск приложения с конфигом: $(CONFIG_PATH)"
	./app

docker-build:
	@echo "Сборка Docker-образа..."
	docker build -t core .

docker-run:
	@echo "Запуск Docker-контейнера..."
	docker run -p 8080:8080 core

migrate:
	@echo "Применение миграций..."
	# Формируем DSN из переменных из .env.
	goose -dir migrations postgres "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" up