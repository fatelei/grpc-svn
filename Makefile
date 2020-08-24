build:
	export GOPROXY=https://goproxy.cn,direct
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0
	go mod tidy

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	protoc -I protos/ protos/grpc_svn.proto protos/common/common.proto --go_out=:internal --go_opt=paths=source_relative --go-grpc_out=:internal --go-grpc_opt=paths=source_relative

server: build proto
	go build -o bin/git_service -i cmd/service/main.go
	./bin/git_service

test:
	go test -v ./pkg/...

lint:
	$(GOPATH)/bin/golangci-lint --skip-dirs=bin -v -e S1023 run
