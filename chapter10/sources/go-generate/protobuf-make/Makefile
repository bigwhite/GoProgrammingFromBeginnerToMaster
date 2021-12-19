all: build

build: gen-protobuf
	go build

gen-protobuf:
	protoc -I ./IDL msg.proto --gofast_out=./msg
