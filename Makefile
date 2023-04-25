clean:
	@rm -rf build
build: clean
	@mkdir build
	@for dir in `ls lambda`; do \
  		GOARCH=amd64 GOOS=linux go build -o build/$$dir lambda/$$dir/main.go; \
  		done