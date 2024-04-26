## help: print usage
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## tidy: format code and tidy modfile
tidy:
	go fmt ./...
	go mod tidy -v

## deploy: deploy function to Google Cloud
deploy:
	gcloud functions deploy hello-http \
		--gen2 \
		--runtime=go122 \
		--region=us-central1 \
		--source=. \
		--entry-point=HelloHTTP \
		--trigger-http \
		--allow-unauthenticated

.PHONY: help tidy deploy
