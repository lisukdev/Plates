clean:
	@rm -f api/*go
	@rm -rf api/.openapi-generator
	@rm -rf build

generate-api: clean
	@openapi-generator generate \
		-i openapi.yaml \
        -g go \
	  	-o ./api \
	  	--additional-properties=isGoSubmodule=true,packageName=api
build: clean generate-api
	@mkdir build
	@for dir in `ls lambda`; do \
  		GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o build/$$dir lambda/$$dir/main.go; \
  		done