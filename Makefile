test:
	go test ./replacer... -v
e2e:
	sh ./test.sh
build:
	go build -o crab ./cmd/...
