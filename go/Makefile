VPATH=cmd:pkg
GOFILES=$(wildcard *.go) $(wildcard **/*.go) $(wildcard */cmd/**/*.go)
YEAR=2018
DAY=1
PADDED_DAY=$(shell printf '%02d' $(DAY))
DAY_EXEC_FILE=$(YEAR)_$(PADDED_DAY)
PKGROOT=github.com/manisenkov/advent-of-code/go

.PHONY: clean build buildgen run rungen test

build: $(GOFILES)
	go build -o bin/$(DAY_EXEC_FILE) $(PKGROOT)/cmd/$(YEAR)/$(PADDED_DAY)

buildgen: $(GOFILES)
	go build -o bin/gensolution $(PKGROOT)/cmd/gensolution

clean:
	rm -rf bin

run: build
	./bin/$(DAY_EXEC_FILE) < inputs/$(YEAR)/$(PADDED_DAY).txt

rungen: buildgen
	./bin/gensolution $(YEAR) $(DAY)

test:
	go test -v ./cmd/$(YEAR)/$(PADDED_DAY)
