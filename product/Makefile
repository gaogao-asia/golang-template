include .env
$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))

OUTPUT ?= $(shell pwd)
SERVICENAME ?= productservice

setup:
    go install google.golang.org/protobuf/cmd/protoc-gen-go; \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc; \
    export PATH="$PATH:$(go env GOPATH)/bin"

gen:
	protoc proto/*.proto --go_out=./ --go-grpc_out=./

clean:
	rm pd/*.go
run:
	go run main.go -config ./config

quickbuild: 
	@echo ""
	@echo "> Build docker image"
	@echo "----------------------------------------"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${OUTPUT}/build2/${SERVICENAME} ./main.go
	docker build --no-cache --platform linux/arm64 -t ${SERVICENAME} -f ${OUTPUT}/build/Dockerfile.quick .
	@echo ""
	@echo "> List docker image"
	@echo "----------------------------------------"
	@docker images | grep ${SERVICENAME}
	@echo ""

run/container:
	@echo "> RUN docker container"
	cp -f ./.env ./build2/ && \
	docker compose -f ${OUTPUT}/build2/docker-compose.yaml up -d
.PHONY: setup gen clean run