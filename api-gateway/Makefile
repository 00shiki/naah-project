protoc:
	@protoc proto/*.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative
	@echo "protobuf generated"

build: protoc
	@go build -o build/app cmd/main.go
	@echo "build successful"

run: build
	@./build/app

clean:
	@rm -rf build/
	@echo "build cleaned"
