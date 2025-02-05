.PHONY: build run docker-build docker-run migrate compose compose-migrate

build:
	@echo "Сборка приложения..."
	go build -o app ./cmd/app

run:
	@echo "Запуск приложения с конфигом: ./configs/local.yaml"
	./app

docker-build:
	@echo "Сборка Docker-образа..."
	docker build -t core .

docker-run:
	@echo "Запуск Docker-контейнера..."
	docker run -p 8080:8080 core

migrate:
	@echo "Применение миграций (локально, если база проброшена на localhost)..."
	goose -dir migrations postgres "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable" up

compose-migrate:
	@echo "Применяем миграции в контейнерной базе..."
	docker run --rm --network core_default core \
      goose -dir migrations postgres "postgres://postgres:postgres@postgres:5432/instagramm_db?sslmode=disable" up

compose:
	@echo "Поднимаем окружение через docker-compose..."
	docker-compose up --build