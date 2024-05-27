SERVICENAME=go_lines_count
SERVICEURL=gitlab.aescorp.ru/dsp_dev/claim/$(SERVICENAME)

FILEMAIN=./cmd/$(SERVICENAME)/main.go
FILEAPP=./bin/$(SERVICENAME)

## build
build:
	clear
	go fmt ./...
	go build -o ./bin/$(SERVICENAME) ./cmd/$(SERVICENAME)/main.go
	cp $(FILEAPP) $(GOPATH)/bin

## run
run:
	clear
	go fmt ./...
	go build -o ./bin/$(SERVICENAME) ./cmd/$(SERVICENAME)/main.go
	cd ./bin/ && ./$(SERVICENAME)

## run.test
run.test:
	clear
	go fmt ./...
	go test -coverprofile cover.out -covermode atomic ./cmd/... ./internal/...
	go tool cover -func=cover.out

## mod
mod:
	clear
	go get -u ./...
	go mod tidy -compat=1.20
	go mod vendor
	go fmt ./...

## lint
lint:
	clear
	go fmt ./...
	golangci-lint run ./cmd/...
	gocyclo -over 10 ./cmd
	gocritic check ./cmd/...
	staticcheck ./cmd/...
	golangci-lint run ./internal/...
	gocyclo -over 10 ./internal
	gocritic check ./internal/...
	staticcheck ./internal/...

## help
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

graph:
	clear
	image_packages ./ ./docs/packages.graphml
conn:
	clear
	image_connections ./cmd/go_lines_count ./docs/connections.graphml $(SERVICENAME)
lines:
	clear
	go_lines_count ./ ./docs/lines_count.txt 10
licenses:
	golicense -out-xlsx=./docs/licenses.xlsx $(FILEAPP)
