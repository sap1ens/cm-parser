.PHONY: build

test:
	go test ./cmd/parser --cover

build:
	go build -o build/parser ./cmd/parser

install:
	go install github.com/sap1ens/cm-parser/cmd/parser