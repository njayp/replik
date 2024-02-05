# runs main.go
.PHONY: run
run:
	go run cmd/server/main.go
	
.PHONY: test
test:
	go test ./...

.PHONY: gen
gen: gen-grpc gen-go

# grpc
.PHONY: gen-grpc
gen-grpc:
	rm -rf ./pkg/api
	mkdir -p ./pkg/api
	protoc \
    	-I=./api \
    	--proto_path=./api \
    	--go_opt=paths=source_relative \
    	--go_out=./pkg/api \
    	--go-grpc_opt=paths=source_relative \
    	--go-grpc_out=./pkg/api \
    	$(shell find ./api -iname "*.proto") 2>&1 > /dev/null

.PHONY: gen-go
gen-go:
	go mod tidy
	go generate ./...

