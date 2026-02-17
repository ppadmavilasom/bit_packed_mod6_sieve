package main

import (
	"math"
	"os"
	"strconv"
)

const (
	// Wheel factorization constants for skipping multiples of 2 and 3.
	wheelSize       = 6 // 2 * 3 = 6
	bitsPerByte     = 8
	unitsPerByte    = 24 // 3 * 8 (skipping multiples of 2 and 3 packs more numbers into bits)
	defaultMaxLimit = 100
)

// Offsets tracks the state for the wheel factorization sieve.
// It skips multiples of 2 and 3, processing only numbers of form 6k+1 and 6k+5.
type Offsets struct {
	start      int
	offset1    int
	offset2    int
	offsetSwap int
	inc        int
	index      int
}

func main() {
	limit := getArg(1, defaultMaxLimit)
	if limit < unitsPerByte {
		limit = unitsPerByte
	}
	bits := calcPrimes(limit)
	if err := os.WriteFile("bits", bits, 0644); err != nil {
		os.Stderr.WriteString("Error writing output file: " + err.Error() + "\n")
		os.Exit(1)
	}
}

// calcPrimes returns a bitset of primes up to n.
// Returns a slice which is efficient (no deep copy of the underlying array).
func calcPrimes(n int) []byte {
	// Initialize offsets for starting at 5
	// alternating increments of 2 and 4 (controlled by bitwise XOR with 6)
	// sequence: 5, 7, 11, 13, 17, 19...
	offsets := Offsets{start: 5, offset1: 7, offset2: 3, offsetSwap: 7 ^ 3, inc: 4}

	// Calculate size of bitset needed
	// We only store bits for numbers coprime to 2 and 3.
	// This mapping effectively packs 24 numbers into one byte (8 bits).
	size := n / unitsPerByte
	bits := make([]byte, size)
	sqrtN := int(math.Sqrt(float64(n)))

	// Initialize bits to 1 (assume all are prime)
	for i := range bits {
		bits[i] = 0xFF
	}

	for ; offsets.start <= sqrtN; offsets.next() {
		// Check if current number is prime (bit is set)
		if bits[offsets.index>>3]&(1<<(7-offsets.index%8)) == 0 {
			continue
		}

		offset := offsets.offset1
		bit := offsets.index + offset
		row := bit >> 3

		// Mark multiples as composite
		for row < size {
			bits[row] &= ^(1 << (7 - (bit % 8)))
			offset ^= offsets.offsetSwap
			bit += offset
			row = bit >> 3
		}
	}

	return bits
}

// next updates the offsets to the next candidate prime (skipping multiples of 2 and 3).
func (o *Offsets) next() {
	o.inc ^= wheelSize
	o.start += o.inc
	o.offset2 += 2
	o.offset1 = (o.start << 1) - o.offset2
	o.index++
	o.offsetSwap = o.offset1 ^ o.offset2
}

// getArg parses the i-th command line argument as an integer, or returns default d.
func getArg(i, d int) int {
	if len(os.Args) > i {
		if n, err := strconv.Atoi(os.Args[i]); err == nil {
			return n
		}
	}
	return d
}
