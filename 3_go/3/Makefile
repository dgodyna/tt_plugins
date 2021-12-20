
prepare:
	mkdir -p solution

build_plugin: prepare
	cd plugin && go build -buildmode=plugin -o ../solution/plugin.so ./cmd/plugin

build_server: prepare
	cd server && go build -overlay="$$(~/go/bin/go-patch-overlay ./patches)" -o ../solution/server ./cmd/server

run_test: build_server build_plugin
	cd solution && ./server

clean:
	rm -rf ./solution