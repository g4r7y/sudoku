PHONY: clean test build

test:
	go test ./...

build: test
	go build -o ./bin/sudoku

clean:
	rm -rf ./bin ./vendor Gopkg.lock


