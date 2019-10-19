.PHONY: build test clean

build:
	@echo "[make] Installing dependencies:"
	go get -v -d ./...
	@echo "[make] Building:"
	go build -v -o klingo

test:
	@echo "[make] Installing test dependencies:"
	go get -v -t -d ./...
	@echo "[make] Running tests:"
	go test -v ./... 

clean:
	rm klingo
