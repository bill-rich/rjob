TEST?=./...

test: test_unit test_integration

test_unit:
	go test $(TEST) --tags=unit

test_integration:
	go test $(TEST) --tags=integration
