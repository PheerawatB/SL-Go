# gen services to server 
protoc --plugin=protoc-gen-go=$(go env GOPATH)/bin/protoc-gen-go \
       --plugin=protoc-gen-go-grpc=$(go env GOPATH)/bin/protoc-gen-go-grpc \
       --go_out=../server --go-grpc_out=../server player.proto


#gen services to client
protoc --plugin=protoc-gen-go=$(go env GOPATH)/bin/protoc-gen-go \
       --plugin=protoc-gen-go-grpc=$(go env GOPATH)/bin/protoc-gen-go-grpc \
       --go_out=../client --go-grpc_out=../client player.proto