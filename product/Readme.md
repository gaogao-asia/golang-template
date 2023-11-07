Step 1: create `proto/processor_message.proto` file

Step 2: install `protobuf`: `brew install protobuf`

Step 1: install go gprc:
```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

