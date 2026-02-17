all: prime

prime: prime.go
	go build -ldflags "-s -w" -o prime prime.go

vet:
	go vet ./prime.go

test:
	go test

clean:
	-rm prime
	-rm bits

.PHONY: all vet test clean
