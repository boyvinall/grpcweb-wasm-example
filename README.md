# gRPC-Web WASM example

Small example server using gRPC-Web with WASM using `go`

This code is an updated copy of <https://github.com/johanbrandhorst/grpcweb-wasm-example> which
generates both CLI and WASM frontends, and separates build logic into targetted make rules.

## Requirements

1. `go` 1.11. See https://golang.org/dl/#go for installation instructions.
1. The Google protobuf compiler, `protoc`.

## Development

Run `make` to regenerate the protofiles and the frontend.

Run `make serve` to start the web server.  Open <https://localhost:10000> in a browser.

Run `make run-cli` to run a compiled CLI client.
