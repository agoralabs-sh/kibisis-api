scripts_dir := ./scripts

all: deploy

deploy:
	doppler run -- doctl serverless deploy "${PWD}"

list:
	doctl serverless functions list

logs:
	doctl serverless activations logs --follow

test:
	$(scripts_dir)/test.sh

watch: deploy
	doppler run -- doctl serverless watch "${PWD}"
