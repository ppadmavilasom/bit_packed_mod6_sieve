FOLDERS=prime prime_optimized utils

all: $(FOLDERS)
$(FOLDERS):
	make -sC $@

vet:
	@for f in $(FOLDERS); do make -sC $$f vet; done

test:
	@for f in $(FOLDERS); do make -sC $$f test; done

clean:
	@for f in $(FOLDERS); do make -sC $$f clean; done

.PHONY: all vet test prime prime_optimized utils
