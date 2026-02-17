# Utils Tool

Helper tool to inspect the binary output of the prime generators.

## Usage

```bash
go run utils.go [command] [file]
```

Commands:
- `count`: Count total set bits (primes).
- `print`: Print bits visually.
- `print_min`: Print bits visually (minimal).

Example:
```bash
go run utils.go count ../bits
```

### How to run utils
`make all` will build a utility called `utils`. If not present, run `make all` or `make utils`

This utility will count primes in the `bits` file if present. Default command is `count` and default file is `bits`

```
$ ./utils count ../prime/bits
23

$ ./utils count bits1
open bits1: no such file or directory
```

### How to print all primes
Assumes `prime` run with no args in ../prime
```
$ ./utils print ../prime/bits
5 7 11 13 17 19 23 
29 31 37 41 43 47 
53 59 61 67 71 73 
79 83 89 97
```

### Get a minimal print
Prints the byte and the starting prime for the byte
```
$ ./utils print_min ../prime/bits
FE 5
DE 29
B7 53
69 79
```
