TEST?=./...

.PHONY: test
test: test_unit test_integration

.PHONY: test_unit
test_unit:
	go test $(TEST) --tags=unit

.PHONY: test_integration
test_integration:
	go test $(TEST) --tags=integration

build:
	mkdir -p bin
	go build -o bin/rclient rclient/main.go
	go build -o bin/rjob rjob/main.go
