VPATH=cmd:pkg
GOFILES=$(wildcard *.go) $(wildcard **/*.go) $(wildcard */solutions/**/*.go)
YEAR=2018
DAY=1
PADDED_DAY=$(shell printf '%02d' $(DAY))
PKGROOT=github.com/manisenkov/advent-of-code

.PHONY: clean build buildgen run rungen test

build: $(GOFILES)
	go build -o bin/main $(PKGROOT)/cmd/main

buildgen: $(GOFILES)
	go build -o bin/gensolution $(PKGROOT)/cmd/gensolution

clean:
	rm -rf bin

run: build
	./bin/main $(YEAR) $(DAY) < inputs/$(YEAR)/$(PADDED_DAY).txt

rungen: buildgen
	./bin/gensolution $(YEAR) $(DAY)

test:
	go test -v ./pkg/$(YEAR)/day$(PADDED_DAY)
