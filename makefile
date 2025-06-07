build:
	go mod tidy
	go build -o bin/app main.go

run:build
	./bin/app

protoc:
	find proto -name "*.proto" -exec protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative {} \;

setup-server:
	go mod init bpay.com/grpc-auth
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	export PATH="$(PATH):$(shell go env GOPATH)/bin"; 
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto; 
	go mod tidy;
	

setup-local:
	go mod init bpay.com/grpc-auth
	export PATH="$(PATH):$(shell go env GOPATH)/bin"; 
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto; 
	go mod tidy; 
	go build -o bin/app main.go; 
	./bin/app
