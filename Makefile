
dev:embed-dev
	@go test
	@go vet
	@golint
	@go build

embed:
	@go-bindata help/...

embed-dev:
	@go-bindata -debug=true help/...

clean:
	@go clean