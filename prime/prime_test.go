package main

import (
	"testing"
)

// Standard Sieve of Eratosthenes for verification
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

	// Reconstruct primes from bits
	// This reverses the logic in calcPrimes to verify what numbers are marked prime
	// Logic:
	// Each byte represents a range of 24 numbers.
	// We skip multiples of 2 and 3.
	// 24 numbers: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24
	// Coprimes to 6 in 1..24: 1, 5, 7, 11, 13, 17, 19, 23 (8 numbers)
	// These 8 numbers correspond to the 8 bits in the byte.

	// Mapping from bit index (0-7) to offset in the 24-block:
	// bit 7 (MSB) -> 1? No, let's look at the code:
	// offsets start at 5.
	// initial offsets: start=5, inc=4 (next is 9? no, inc^=6 so 4, 2, 4, 2...)
	// Sequence: 5, 7 (5+2), 11 (7+4), 13 (11+2), 17 (13+4), 19 (17+2), 23 (19+4), 25 (23+2)...

	// Wait, the code initialization:
	// Offsets{start: 5, ..., inc: 4}
	// next(): inc ^= 6 (becomes 2). start += 2 -> 7.
	// next(): inc ^= 6 (becomes 4). start += 4 -> 11.
	// So sequence is 5, 7, 11, 13, 17, 19, 23...

	// These correspond to bits.
	// byte 0 covers 0..23?
	// "bits" array index `i`.
	// For a given `i`, it covers range `[i*24, (i+1)*24)`.
	// Inside that range, the bits correspond to the sequence 5, 7, 11, 13, 17, 19, 23, 25(next block 1)...
	// Actually, let's look at the bit index logic.
	// `b[o.index >> 3]` accessing byte.
	// `o.index` starts at 0? No, struct default is 0.
	// `o.index` increments by 1 for each step in the loop (5, 7, 11...)
	// So index 0 -> 5
	// index 1 -> 7
	// index 2 -> 11
	// index 3 -> 13
	// ...
	// index 7 -> 23
	// index 8 -> 25 (which is 1 mod 24 of the next block)

	primesToCheck := sieveOfEratosthenes(limit)

	// We need to verify that for every prime P (where P > 3), the corresponding bit is 1.
	// And for every composite C (where C % 2 != 0 && C % 3 != 0), the corresponding bit is 0.

	// Map index to number
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
