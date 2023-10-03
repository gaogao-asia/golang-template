include .env
$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))

OUTPUT ?= $(shell pwd)
SERVICENAME ?= gotemplate

init:
	docker compose -f starter/docker-compose.yaml up -d
	go install github.com/vektra/mockery/v2@latest
	go install github.com/google/wire/cmd/wire@latest
	if ! docker network inspect apm >/dev/null 2>&1; then \
        docker network create apm; \
    fi
	if ! docker plugin inspect loki &> /dev/null; then \
  		docker plugin install grafana/loki-docker-driver:latest --alias loki --grant-all-permissions; \
	fi
	docker-compose -f ./starter/monitor/docker-compose.yaml up --build -d 
	docker image prune --force
	docker system prune --volumes --force

init/down:
	docker compose -f starter/docker-compose.yaml down
	docker-compose -f ./starter/monitor/docker-compose.yaml down
	docker image prune --force
	docker system prune --volumes --force
	
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

di:
	wire ./internal/di/...

quickbuild: 
	@echo ""
	@echo "> Build docker image"
	@echo "----------------------------------------"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${OUTPUT}/build/${SERVICENAME} ./cmd/*.go
	docker build --no-cache --platform linux/arm64 -t ${SERVICENAME} -f ${OUTPUT}/build/Dockerfile.quick .
	@echo ""
	@echo "> List docker image"
	@echo "----------------------------------------"
	@docker images | grep ${SERVICENAME}
	@echo ""

run/container:
	@echo "> RUN docker container"
	cp -f ./.env ./build/ && \
	docker compose -f ${OUTPUT}/build/docker-compose.yaml up -d

.PHONY: init init/down run mock unit-test di
