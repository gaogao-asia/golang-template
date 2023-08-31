include .env
$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))

init:
	docker compose -f starter/docker-compose.yaml up -d
init/down:
	docker compose -f starter/docker-compose.yaml down
run:
	go run cmd/main.go
