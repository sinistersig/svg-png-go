IMG=svg-png
OUTPUT_DIR=/Users/sigi
DOCKER_RUN_CMD=docker run --rm -it -v $(OUTPUT_DIR):/home $(IMG)

docker-pull:
	docker pull sinistersig/alpine-go:latest

docker-build:
	docker build -t $(IMG) .

docker-run:
	$(DOCKER_RUN_CMD) /bin/bash "run-script.sh"

docker-test:
	$(DOCKER_RUN_CMD) go "test"

