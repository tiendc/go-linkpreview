build:
	@go build -v ./...

test:
	@go test -cover  -v ./...

mod:
	go mod tidy && go mod vendor
