PACKAGES=$(shell go list ./... | grep -v amigo/vendor)

all:

test:
	go vet $(PACKAGES)
	go test $(PACKAGES)

deploy:
	gcloud app deploy

gae-logs:
	gcloud app logs tail -s default

gae-browse:
	gcloud app browse

.PHONY: test deploy gae-logs gae-browse