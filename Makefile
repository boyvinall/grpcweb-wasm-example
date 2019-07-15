GOPATH=$(realpath ../../../..)
GOROOT=$(shell go env GOROOT)

.PHONY: all
all: out/server out/cli

.PHONY: clean
clean:
	rm -rf ./frontend/html ./proto/web.pb.go ./frontend/bundle ./out ./vendor

frontend/html/index.html: $(GOROOT)/misc/wasm/wasm_exec.html
	cp $< $@
	sed -i -e 's;</button>;</button>\n\t<div id=\"target\"></div>;' $@

frontend/html/wasm_exec.js: $(GOROOT)/misc/wasm/wasm_exec.js
	cp $< $@

proto/web.pb.go: proto/web.proto
	protoc -I. ./proto/web.proto --go_out=plugins=grpc:$(GOPATH)/src

frontend/html/test.wasm: proto/web.pb.go vendor
	GOOS=js GOARCH=wasm go build -o $@ ./frontend 

.PHONY: out/server
out/server: proto/web.pb.go frontend/bundle/bundle.go vendor
	go build -o $@

vendor:
	dep ensure -v

frontend/bundle/bundle.go: frontend/html/test.wasm frontend/html/index.html frontend/html/wasm_exec.js
	mkdir -p $(dir $@)
	cd frontend && go run ./assets_generate.go

.PHONY: out/cli
out/cli: proto/web.pb.go vendor
	go build -o $@ ./frontend

.PHONY: serve
serve: out/server
	out/server
