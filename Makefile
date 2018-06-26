default: generate-mocks

.PHONEY: generate-mocks
generate-mocks:
	mkdir -p ./internal/app/mocks/
	go generate ./...
