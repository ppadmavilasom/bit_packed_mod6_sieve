package main

import (
	"math"
	"os"
	"strconv"
)

const (
	TWO     = 2
	SIX     = 6
	UNIT    = 24
	DEFAULT = 100
)

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
	n := getArg(1, 1, DEFAULT)
	if n < UNIT {
		n = UNIT
	}
	calc_primes(n)
}

func calc_primes(n int) {
	o := Offsets{
		start:   5,
		offset1: 7, offset2: 3,
		offsetSwap: 7 ^ 3,
		inc:        4,
		squareSeq1: 1, squareSeq2: 1, squareInc: 1, squareSeq: 2,
	}
	l := n / UNIT
	b := make([]byte, l)
	sqrtn := int(math.Sqrt(float64(n)))
	for i := range b {
		b[i] = 0xFF
	}
	for ; o.start <= sqrtn; o.next() {
		if b[o.index>>3]&(1<<(7-o.index%8)) == 0 {
			continue
		}
		offset := o.offset1
		if (o.index & 1) == 1 {
			offset = o.offset2
		}
		bit := ((o.squareIndex + 1) << 3) - 1
		row := bit >> 3
		for row < l {
			b[row] &= ^(1 << (7 - (bit % 8)))
			offset ^= o.offsetSwap
			bit += offset
			row = bit >> 3
		}
	}
	os.WriteFile("bits", b, 0644)
}

func (o *Offsets) next() {
	o.inc ^= SIX
	o.start += o.inc
	o.offset2 += 2
	o.offset1 = (o.start << 1) - o.offset2
	o.offsetSwap = o.offset1 ^ o.offset2
	o.index++
	o.squareIndex += o.squareInc
	if o.squareSeq == 2 {
		o.squareSeq2 += 2
		o.squareInc = o.squareSeq2
	} else {
		o.squareSeq1++
		o.squareInc = o.squareSeq1
	}
	o.squareSeq ^= TWO
}

func getArg(i, p, d int) int {
	if len(os.Args) > p {
		n, _ := strconv.Atoi(os.Args[i])
		return n
	}
	return d
}
