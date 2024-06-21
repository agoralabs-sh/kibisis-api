scripts_dir := ./scripts

.PHONY: services

all: setup swagger deploy

deploy:
	doppler run -- doctl serverless deploy "${PWD}"

list:
	doctl serverless functions list

logs:
	doctl serverless activations logs --follow

services:
	doppler run -- docker compose -f ./deployments/docker-compose.yml up --build

services-relayer:
	doppler run -- docker compose -f ./deployments/docker-compose.yml up relayer --build

setup:
	go install github.com/swaggo/swag/cmd/swag@latest # install swag cli

swagger:
	$(scripts_dir)/create_swagger_docs.sh 1 # v1

test:
	$(scripts_dir)/test.sh

watch: deploy
	doppler run -- doctl serverless watch "${PWD}"
