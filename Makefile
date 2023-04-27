clean:
	@rm -f api/*go
	@rm -rf api/.openapi-generator
	@rm -rf build

generate-api:
	@docker run --rm \
		-v ${PWD}:/local openapitools/openapi-generator-cli generate \
		-i /local/openapi.yaml \
		-g go \
		-c /local/openapi-generator-config.json \
		-o /local/api

test: generate-api
	go test ./...

build: clean generate-api
	@mkdir build
	@for dir in `ls lambda`; do \
  		GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o build/$$dir lambda/$$dir/main.go; \
  		done
