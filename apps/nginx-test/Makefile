.PHONY: dev start stop

dev:
	docker compose --file ./deployments/compose.yml watch

start:
	docker compose --file ./deployments/compose.yml up --detach

stop:
	docker compose --file ./deployments/compose.yml down
