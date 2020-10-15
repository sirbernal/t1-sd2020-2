grpc:
	export GO111MODULE=on  
	go get github.com/golang/protobuf/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0
	export PATH="$PATH:$(go env GOPATH)/bin"

cliente_logistica: grpc
	protoc --proto_path=. --go_out=plugins=grpc:proto proto/cliente_logistica.proto
	protoc --proto_path=. --go_out=plugins=grpc:proto proto/camion_logistica.proto
print:
	echo "print"
	