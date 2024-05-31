scripts_dir := ./scripts

.PHONY: deploy

deploy:
	doctl serverless deploy "${PWD}"

list:
	doctl serverless functions list

logs:
	doctl serverless activations logs --follow

watch:
	doctl serverless watch "${PWD}"
