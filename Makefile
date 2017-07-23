PACKAGES=$(shell go list ./... | grep -v amigo/vendor)

all: build

.PHONY: test
test:
	go vet $(PACKAGES)
	go test $(PACKAGES)

.PHONY: build
build:
	go build .

.PHONY: deploy
deploy:
	gcloud app deploy

.PHONY: gae-logs
gae-logs:
	gcloud app logs tail -s default

.PHONY: gae-browse
gae-browse:
	gcloud app browse
