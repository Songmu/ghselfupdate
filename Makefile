deps:
	go get -d -v .

test-deps:
	go get -d -v -t .

devel-deps: test-deps
	go get ${u} github.com/golang/lint/golint

test: deps
	go test

lint: devel-deps
	go vet
	golint -set_exit_status

.PHONY: deps test-deps devel-deps test lint
