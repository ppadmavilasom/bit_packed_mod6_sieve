package main

import (
	"math"
	"os"
	"strconv"
)

const (
	// Wheel factorization constants
	wheelSize       = 6
	wheelToggle     = 2 // Used to toggle square increment?
	bitsPerByte     = 8
	unitsPerByte    = 24
	defaultMaxLimit = 100
)

// Offsets tracks the state for the wheel factorization sieve.
// This version includes additional state for "square" optimization.
type Offsets struct {
	start       int
	offset1     int
	offset2     int
	offsetSwap  int
	inc         int
	index       int
	squareIndex int
	squareSeq1  int
	squareSeq2  int
	squareInc   int
	squareSeq   int
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
	// Initialize offsets
	offsets := Offsets{
		start:      5,
		offset1:    7,
		offset2:    3,
		offsetSwap: 7 ^ 3,
		inc:        4,
		squareSeq1: 1,
		squareSeq2: 1,
		squareInc:  1,
		squareSeq:  2,
	}

	size := n / unitsPerByte
	bits := make([]byte, size)
	sqrtN := int(math.Sqrt(float64(n)))

	for i := range bits {
		bits[i] = 0xFF
	}

	for ; offsets.start <= sqrtN; offsets.next() {
		if bits[offsets.index>>3]&(1<<(7-offsets.index%8)) == 0 {
			continue
		}

		// Optimized inner loop logic
		offset := offsets.offset1
		if (offsets.index & 1) == 1 {
			offset = offsets.offset2
		}

		// Calculate starting bit position based on square optimization
		bit := ((offsets.squareIndex + 1) << 3) - 1
		row := bit >> 3

		for row < size {
			bits[row] &= ^(1 << (7 - (bit % 8)))
			offset ^= offsets.offsetSwap
			bit += offset
			row = bit >> 3
		}
	}

	return bits
}

// next updates the offsets to the next candidate prime.
func (o *Offsets) next() {
	o.inc ^= wheelSize
	o.start += o.inc
	o.offset2 += 2
	o.offset1 = (o.start << 1) - o.offset2
	o.offsetSwap = o.offset1 ^ o.offset2
	o.index++

	// Square optimization updates
	o.squareIndex += o.squareInc
	if o.squareSeq == 2 {
		o.squareSeq2 += 2
		o.squareInc = o.squareSeq2
	} else {
		o.squareSeq1++
		o.squareInc = o.squareSeq1
	}
	o.squareSeq ^= wheelToggle
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
