
compile_proto:
	cd server && protoc --plugin=protoc-gen-go=/Users/dgodyna/go/bin/protoc-gen-go -Iapi --go_out=module=tt_go_plugin_server:. api/*.proto
	cd server && protoc --plugin=protoc-gen-go-grpc=/Users/dgodyna/go/bin/protoc-gen-go-grpc -Iapi --go-grpc_out==module=tt_go_plugin_server:. api/*.proto

prepare:
	mkdir -p solution

build_plugin: prepare
	cd plugin && go build -o ../solution/plugin ./cmd/plugin

build_server: prepare
	cd server && go build -o ../solution/server ./cmd/server

run: build_plugin build_server
	cd solution && ./server

clean:
	rm -rf ./solution