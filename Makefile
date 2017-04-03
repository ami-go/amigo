PACKAGES=$(shell go list ./... | grep -v amigo/vendor)

all:

test:
	go vet $(PACKAGES)
	go test $(PACKAGES)

deploy:
	gcloud app deploy

.PHONY: test deploy