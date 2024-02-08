.PHONY: dev start stop

dev:
	docker compose --file ./deployments/docker-compose.yml watch

start:
	docker compose --file ./deployments/docker-compose.yml up --detach

stop:
	docker compose --file ./deployments/docker-compose.yml down
