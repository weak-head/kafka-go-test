PRODUCER_IMAGE_ID ?= github.com/weak-head/go-kafka-test/producer
PRODUCER_IMAGE_TAG ?= 0.0.1

CONSUMER_IMAGE_ID ?= github.com/weak-head/go-kafka-test/consumer
CONSUMER_IMAGE_TAG ?= 0.0.1


.PHONY: build-docker-producer
build-docker-producer:
	cd producer && docker build -t $(PRODUCER_IMAGE_ID):$(PRODUCER_IMAGE_TAG) .


.PHONY: build-docker-consumer
build-docker-consumer:
	cd consumer && docker build -t $(CONSUMER_IMAGE_ID):$(CONSUMER_IMAGE_TAG) .


.PHONY: build-docker
build-docker: build-docker-producer build-docker-consumer


.PHONY: load-images
load-images:
	kind load docker-image $(PRODUCER_IMAGE_ID):$(PRODUCER_IMAGE_TAG)
	kind load docker-image $(CONSUMER_IMAGE_ID):$(CONSUMER_IMAGE_TAG)


.PHONY: deploy
deploy: build-docker load-images
	kubectl create -f deployment/producer-p1-deployment.yaml
	kubectl create -f deployment/producer-p2-deployment.yaml
	kubectl create -f deployment/consumer-g1-1-deployment.yaml
	kubectl create -f deployment/consumer-g1-2-deployment.yaml
	kubectl create -f deployment/consumer-g2-1-deployment.yaml


.PHONY: clean
clean:
	kubectl delete -f deployment/producer-p1-deployment.yaml
	kubectl delete -f deployment/producer-p2-deployment.yaml
	kubectl delete -f deployment/consumer-g1-1-deployment.yaml
	kubectl delete -f deployment/consumer-g1-2-deployment.yaml
	kubectl delete -f deployment/consumer-g2-1-deployment.yaml