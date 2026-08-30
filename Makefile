include .env
export

export PROJECT_ROOT=$(shell pwd)
export HOST_UID := $(shell id -u)
export HOST_GID := $(shell id -g)



env-up:
	@if [ ! -d ${PROJECT_ROOT}/out/pgdata ]; then \
		mkdir ${PROJECT_ROOT}/out/pgdata; \
	fi
	docker compose up -d todoapp-postgres

env-down:
	@docker compose down todoapp-postgres

env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность утери данных. [y/n]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres && \
		docker compose down port-forwarder && \
		rm -rf ${PROJECT_ROOT}/out/pgdata/ && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi



env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder



migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует seq";\
		exit 1; \
	fi;
	docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует action";\
		exit 1; \
	fi;
	docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		$(action)

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down



todoapp-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/todoapp/main.go