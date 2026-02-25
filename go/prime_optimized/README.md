# Optimized Prime Number Generator

This is a high-performance implementation of the Sieve of Eratosthenes using a wheel factorization optimization (skipping multiples of 2 and 3) and an additional square calculation (logic for bit start position) for speed.

## Usage

Run the program with a maximum limit `N`.

```bash
go run prime_optimized.go [limit]
```

## Output

The program outputs a binary file named `bits`.
Each byte represents 8 numbers potentially coprime to 6 within a range of 24 integers.
