all: prime_optimized

prime_optimized: prime_optimized.c
	cc -Ofast prime_optimized.c -lm -o prime_optimized

clean:
	-rm prime_optimized

.PHONY: all clean
.SILENT:
