FOLDERS=prime prime_optimized utils

all: $(FOLDERS)
$(FOLDERS):
	make -sC $@

vet:
	@for f in $(FOLDERS); do make -sC $$f vet; done

test:
	@for f in $(FOLDERS); do make -sC $$f test; done

integration: all
	./check_shasums.sh ./prime/prime 8
	./check_shasums.sh ./prime_optimized/prime_optimized 8

clean:
	@for f in $(FOLDERS); do make -sC $$f clean; done

.PHONY: all vet test integration prime prime_optimized utils
.SILENT:
