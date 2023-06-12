build:
	go build -o ./bin/go-another-blockchain-node -v

run: build
	./bin/go-another-blockchain-node

test:
	go test -v ./...