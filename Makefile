.PHONY: vendor
vendor:
	go mod vendor

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: download
download:
	go mod download

.PHONY: build
build:
	make build-service && make build-migration

.PHONY: build-service
build-service:
	go build -mod=vendor -o .\build\bin\main .\cmd\bin\main.go

.PHONY: build-migration
build-migration:
	go build -mod=vendor -o .\build\bin\migration .\cmd\bin\main.go

.PHONY: run
run:
	go run -mod=vendor .\cmd\reminder\main.go

.PHONY: test
test:
	go test -mod=vendor .\cmd\reminder\main.go

.PHONY: graph
graph:
	go mod graph
