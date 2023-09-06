include .env
$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))

init:
	docker compose -f starter/docker-compose.yaml up -d
	go install github.com/vektra/mockery/v3@latest
init/down:
	docker compose -f starter/docker-compose.yaml down
run:
	go run cmd/main.go

mock: 
	go generate ./...

unit-test:
	@echo ""
	@echo "--------- Run unittest in internal module ----------"
	@echo ""
	go test -cover ./internal/... -count=1

.PHONY: init init/down run mock unit-test
