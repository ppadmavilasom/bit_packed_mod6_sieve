FOLDERS=c/prime c/prime_optimized go/prime go/prime_optimized utils
GO_FOLDERS=go/prime go/prime_optimized utils

all:
	@for f in $(FOLDERS); do make -sC $$f all; done

vet:
	@for f in $(GO_FOLDERS); do make -sC $$f vet; done

test:
	@for f in $(GO_FOLDERS); do make -sC $$f test; done

integration: all
	./check_shasums.sh ./go/prime/prime 8
	./check_shasums.sh ./go/prime_optimized/prime_optimized 8
	./check_shasums.sh ./c/prime/prime 8
	./check_shasums.sh ./c/prime_optimized/prime_optimized 8

clean:
	@for f in $(FOLDERS); do make -sC $$f clean; done

.PHONY: all vet test integration utils
.SILENT:
