PACKAGES=$(shell go list ./... | grep -v amigo/vendor)

all:

test:
	go vet $(PACKAGES)
	go test $(PACKAGES)

.PHONY: test