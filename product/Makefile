install:
    @go install google.golang.org/protobuf/cmd/protoc-gen-go
    @go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
    @export PATH="$PATH:$(go env GOPATH)/bin"

gen:
	protoc proto/*.proto --go_out=./ --go-grpc_out=./

clean:
	rm pd/*.go
run:
	go run main.go