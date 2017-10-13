.DEFAULT_TARGET := build-dev
IMAGE_NAME := agrea/go-test-runner
VERSION := "0.2"

.PHONY: build-dev
build-dev:
	$(info --- Building dev image)
	docker build -t $(IMAGE_NAME):dev .

.PHONY: build-release
build-release:
	$(info --- Building release image)
	docker build -t $(IMAGE_NAME):latest -t $(IMAGE_NAME):$(VERSION) .

.PHONY: push
push:
	$(info --- Pushing to Docker Hub)
	docker push $(IMAGE_NAME):latest
	docker push $(IMAGE_NAME):$(VERSION)

.PHONY: release
release: build-release push
	$(info --- Building release image)
	docker build -t $(IMAGE_NAME):latest -t $(IMAGE_NAME):$(VERSION) .
