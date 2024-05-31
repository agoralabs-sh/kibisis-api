scripts_dir := ./scripts

all: setup deploy

deploy:
	doctl serverless deploy "${PWD}"

list:
	doctl serverless functions list

logs:
	doctl serverless activations logs --follow

setup:
	$(scripts_dir)/setup.sh

watch: deploy
	doctl serverless watch "${PWD}"
