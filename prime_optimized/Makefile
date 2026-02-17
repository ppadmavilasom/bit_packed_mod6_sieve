all: prime_optimized

prime_optimized: prime_optimized.go
	go build -ldflags "-s -w" -o prime_optimized prime_optimized.go

vet:
	go vet ./prime_optimized.go

test:
	go test

clean:
	-rm prime_optimized
	-rm bits

.PHONY: all vet test clean
