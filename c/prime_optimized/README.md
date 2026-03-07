# Optimized Prime Number Generator

This is a high-performance implementation of the Sieve of Eratosthenes using a wheel factorization optimization (skipping multiples of 2 and 3) and an additional square calculation (logic for bit start position) for speed.

## How to build
Run `make` to build binary. There will be a file `./prime_optimized` in this directory.

## Usage

Run the program with a maximum limit `N`.


```bash
./prime_optimized [limit]
```

Example:
```bash
./prime_optimized $((10**8))
ls -alh bits

-rw-rw-r-- 1 pgp pgp 4.0M Mar  6 19:31 bits
```

## Output

The program outputs a binary file named `bits`.
Each byte represents 8 numbers potentially coprime to 6 within a range of 24 integers.
