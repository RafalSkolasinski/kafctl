LATEST_TAG ?= latest

IMAGE_NAME ?= rafalskolasinski/kafctl
IMAGE_VERSION ?= ${LATEST_TAG}

install:
	go install .

build:
	go build .

docker-build:
	docker build -t ${IMAGE_NAME}:${LATEST_TAG} .

docker-push:
	docker push ${IMAGE_NAME}:${LATEST_TAG}
