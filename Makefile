APP := security-policy-admissions
VERSION ?= "latest"
IMAGE := ghcr.io/jspc/$(APP):$(VERSION)

BINARY := app

.PHONY: default
default: $(BINARY)

$(BINARY): *.go go.*
	CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o $@ && upx $@

.PHONY: docker-build docker-push
docker-build: $(BINARY)
	docker build -t $(IMAGE) .

docker-push:
	docker push $(IMAGE)
