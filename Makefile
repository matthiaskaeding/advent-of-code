.PHONY: run
run:
	go run main.go 

.PHONY: testv
testv:
	go test -v ./...

.PHONY: test
test:
	go test ./...

check:
	staticcheck ./...
