TEST?=./...

.PHONY: test
test: test_unit test_integration

.PHONY: test_unit
test_unit:
	go test $(TEST) --tags=unit

.PHONY: test_integration
test_integration:
	go test $(TEST) --tags=integration
