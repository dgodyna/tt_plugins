prepare:
	mkdir -p solution

build_plugin:
	cd plugin && go build -buildmode=c-shared -o ../solution/plugin.so ./cmd/plugin

build_server:
	cd server && go build -o ../solution/server ./cmd/server

run: 
	cd solution && LD_LIBRARY_PATH=/Users/dgodyna/work/My/tt_plugins/3_go/4/solution ./server