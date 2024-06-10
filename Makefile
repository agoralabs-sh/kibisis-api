scripts_dir := ./scripts

all: setup swagger deploy

deploy:
	doppler run -- doctl serverless deploy "${PWD}" --verbose

list:
	doctl serverless functions list

logs:
	doctl serverless activations logs --follow

setup:
	go install github.com/swaggo/swag/cmd/swag@latest # install swag cli

swagger:
	$(scripts_dir)/create_swagger_docs.sh

test:
	$(scripts_dir)/test.sh

watch: deploy
	doppler run -- doctl serverless watch "${PWD}"
