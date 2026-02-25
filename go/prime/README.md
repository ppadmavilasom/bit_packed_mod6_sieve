# Prime Number Generator

This is a high-performance implementation of the Sieve of Eratosthenes using a wheel factorization optimization (skipping multiples of 2 and 3).

## Usage

Run the program with a maximum limit `N` to generate primes up to `N`.

```bash
go run prime.go [limit]
```

Example:

```bash
go run prime.go 1000
```

## Output

The program outputs a binary file named `bits`. The bits in this file represent the primality of numbers in the sequence (excluding multiples of 2 and 3).

- A set bit (1) indicates a prime number.
- A cleared bit (0) indicates a composite number.

Each byte represents a range of 24 numbers. The mapping is optimized to store only candidates coprime to 6.
