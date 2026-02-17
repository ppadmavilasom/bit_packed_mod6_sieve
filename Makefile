all: prime prime_optimized utils

prime: prime.go
	go build -ldflags "-s -w" -o prime prime.go

prime_optimized: prime_optimized.go
	go build -ldflags "-s -w" -o prime_optimized prime_optimized.go

utils: utils.go
	go build -ldflags "-s -w" -o utils utils.go

vet:
	go vet ./prime.go
	go vet ./prime_optimized.go
	go vet ./utils.go

test:
	go run ./prime.go

.PHONY: all utils vet test
