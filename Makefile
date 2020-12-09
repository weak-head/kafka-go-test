PRODUCER_IMAGE_ID ?= github.com/weak-head/go-kafka-test/producer
PRODUCER_IMAGE_TAG ?= 0.0.1

.PHONY: build-docker-producer
build-docker-producer:
	cd producer && docker build -t $(PRODUCER_IMAGE_ID):$(PRODUCER_IMAGE_TAG) .