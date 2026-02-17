all: utils

utils: utils.go
	go build -ldflags "-s -w" -o utils utils.go

vet:
	go vet ./utils.go

test:

clean:
	-rm utils
	-rm bits

.PHONY: all vet clean
