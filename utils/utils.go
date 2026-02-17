package main

import (
	"fmt"
	"math/bits"
	"os"
	"syscall"
)

var bitCounts [256]int

func init() {
	for i := 0; i < 256; i++ {
		bitCounts[i] = bits.OnesCount8(uint8(i))
	}
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	file := "bits"
	if len(os.Args) > 2 {
		file = os.Args[2]
	}

	var err error
	switch command {
	case "print":
		err = processFile(file, false, false)
	case "print_min":
		err = processFile(file, false, true)
	case "count":
		err = processFile(file, true, false)
	default:
		printUsage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: utils [command] [file]")
	fmt.Println("Commands:")
	fmt.Println("  count      Count total set bits")
	fmt.Println("  print      Print bits visually")
	fmt.Println("  print_min  Print bits visually (minimal)")
	fmt.Println("Default file is 'bits'")
}

func processFile(filename string, countOnly, minOutput bool) error {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat file: %w", err)
	}

	if fi.Size() == 0 {
		return nil
	}

	// Use mmap for efficient reading of large files
	fd := int(f.Fd())
	data, err := syscall.Mmap(fd, 0, int(fi.Size()), syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		return fmt.Errorf("mmap failed: %w", err)
	}
	defer syscall.Munmap(data)

	if countOnly {
		total := 0
		for _, b := range data {
			total += bitCounts[b]
		}
		fmt.Println(total)
	} else {
		if minOutput {
			printArray(data, printByteMin)
		} else {
			printArray(data, printByte)
		}
	}
	return nil
}

func printArray(data []byte, printFunc func(int, byte)) {
	start := 5
	for _, b := range data {
		printFunc(start, b)
		start += 24
	}
}

func printByte(start int, b byte) {
	skip := 0
	for i := 7; i >= 0; i-- {
		if b&(1<<i) > 0 {
			fmt.Print(start, " ")
		}
		start += (2 + skip)
		// Toggle skip between 0 and 2 (wheel increments)
		skip ^= 2
	}
	fmt.Println()
}

func printByteMin(start int, b byte) {
	if b == 0 {
		fmt.Println()
		return
	}
	skip := 0
	fmt.Printf("%.2X ", b)
	for i := 7; i >= 0; i-- {
		if b&(1<<i) > 0 {
			fmt.Println(start)
			break
		}
		start += (2 + skip)
		// Toggle skip between 0 and 2 (wheel increments)
		skip ^= 2
	}
}
