build_plugin:
	gcc -shared -o libplugin.so -fPIC plugin.c

build_server:
	gcc server.c -o server

run_server: build_plugin build_server
	LD_LIBRARY_PATH=$(CURDIR) ./server 'hello gophers!'

clean:
	rm -rf ./*.so
	rm -rf ./server




