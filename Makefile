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

run cliente:
	cd cliente
	go run cliente.go	

run logistica:
	go run logistica/logistica.go

run camion:
	go run camion/camion.go		

run financiero:
	go run financiero/financiero.go	
