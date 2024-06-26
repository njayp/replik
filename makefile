# runs server
.PHONY: run
run:
	go run ./cmd/server/main.go

.PHONY: build
build:
	rm -rf ./output
	mkdir -p ./output
	go build -o ./output/replik ./cmd/cli/main.go
	go build -o ./output/server ./cmd/server/main.go
	
.PHONY: test
test:
# -timeout 10m
	go test -v ./...

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
	go get -u ./...
	go mod tidy
	go generate ./...
	