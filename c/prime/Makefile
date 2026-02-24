all: prime

prime: prime.c
	cc -Ofast prime.c -lm -o prime

valgrind: prime
	valgrind --leak-check=full ./prime 1000

clean:
	-rm prime

.PHONY: all clean valgrind
.SILENT:
