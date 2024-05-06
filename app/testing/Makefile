## help: print usage
help:
	@echo 'Usage: ' && sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## test/browser: run all browser tests
test/browser:
	@npx playwright install-deps
	cd browser && npx playwright test

## test/load: run all load tests
test/load:
	@go install go.k6.io/k6@latest
	cd load && k6 run home.js
