clean:
	@rm -f api/*go
	@rm -rf api/.openapi-generator
	@rm -rf build

generate-api: clean
	@docker run --rm \
		-v ${PWD}:/local openapitools/openapi-generator-cli generate \
		-i /local/openapi.yaml \
		-g go \
		-c /local/openapi-generator-config.json \
		-o /local/api

test: clean generate-api
	go test ./...

build: clean generate-api test
	@mkdir build
	@for dir in `ls lambda`; do \
  		GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o build/$$dir lambda/$$dir/main.go; \
  		done

deploy: build
	terraform -chdir=terraform apply -auto-approve -input=false
