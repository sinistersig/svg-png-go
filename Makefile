IMG=receiptful-test
DOCKER_RUN_CMD=docker run --rm -it $(IMG)

docker-pull:
	docker pull sinistersig/alpine-go:latest

docker-build:
	docker build -t $(IMG) .

docker-run:
	$(DOCKER_RUN_CMD) /bin/bash "run-script.sh"

docker-test:
	$(DOCKER_RUN_CMD) go "test"

