MAJOR   := 0
MINOR   := 0
PATCH   := 0
VERSION := $(MAJOR).$(MINOR).$(PATCH)

default:
	go build

install:
	go install

clean:
	go clean

version:
	@echo v$(VERSION)
