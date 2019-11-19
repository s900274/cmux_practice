export GO111MODULE=on

all: clean bld

bld:
	go build -o init/cmux-server ./cmd/server
	#go build -o init/http-client ./cmd/http-client
	go build -o init/grpc-client ./cmd/grpc-client

clean:
	@rm -f init/cmux-server
	#@rm -f init/http-client
	@rm -f init/grpc-client

cleanlog:
	@rm -f log/*log*