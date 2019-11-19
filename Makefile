export GO111MODULE=on

all: clean bld

bld:
	go build -o init/cmux-server ./cmd/server
	go build -o init/cmux-client ./cmd/client

clean:
	@rm -f init/cmux-server
	@rm -f init/client

cleanlog:
	@rm -f log/*log*