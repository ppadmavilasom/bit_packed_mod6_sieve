package main

import (
	"testing"
)

// Standard Sieve for verification
func sieveOfEratosthenes(n int) []bool {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}
	return isPrime
}

func TestCalcPrimes(t *testing.T) {
	limit := 1000
	content := calcPrimes(limit)

	// Reconstruct primes logic
	// Each byte represents a range of 24 numbers.
	// We skip multiples of 2 and 3.
	// Sequence in block of 24: 1, 5, 7, 11, 13, 17, 19, 23 (8 numbers)
	// These correspond to bits 7, 6, 5, 4, 3, 2, 1, 0 respectively.

	primesToCheck := sieveOfEratosthenes(limit)

	wheel := []int{2, 4}
	val := 5
	wIdx := 0

	for idx := 0; idx < len(content)*8; idx++ {
		byteVal := content[idx/8]
		bitPos := 7 - (idx % 8)
		isSet := (byteVal & (1 << bitPos)) != 0

		if val >= limit {
			break
		}

		if primesToCheck[val] {
			if !isSet {
				t.Errorf("Number %d is prime but bit was 0", val)
			}
		} else {
			if isSet {
				t.Errorf("Number %d is composite but bit was 1", val)
			}
		}

		val += wheel[wIdx]
		wIdx = (wIdx + 1) % 2
	}
}
