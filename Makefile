COVERAGE_FILE:=sub.cov

build:
	go build -o ../../bin/covtest ./main.go

test:
	go test ./...

instrumented:
	go test -o ../../bin/covtest.coverage -c -covermode=count -coverpkg ./...

run:
	 ../../bin/covtest.coverage -test.coverprofile=$(COVERAGE_FILE) run_instrumented_binary

check:
	curl -X GET localhost:8080/two
	curl -X GET localhost:8080/text

show-coverage:
	go tool cover -html=$(COVERAGE_FILE)

.PHONY: build test instrumented run check show-coverage