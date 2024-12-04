tidy:
	go mod tidy

.PHONY: test
test:
	go test --cover ./...