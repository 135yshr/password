.PHONY: test
test:
	go test -v -parallel 2 ./...

.PHONY: clean
clean:
	@echo "clean"
