include .env
$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))

init:
	docker compose -f starter/docker-compose.yaml up -d
	go install github.com/vektra/mockery/v2@latest
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

integration-test/db/up:
	cp -f ./starter/init.sql ./integration_test/init.sql && \
	docker-compose -f ./integration_test/docker-compose.yaml up -d
integration-test/db/down:
	docker-compose -f ./integration_test/docker-compose.yaml down

integration-test:
	@echo ""
	@echo "--------- Run integration test  ----------"
	@echo ""
	go test ./integration_test/...  -failfast  -count=1

.PHONY: init init/down run mock unit-test
