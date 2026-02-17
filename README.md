# bit_packed_mod6_sieve
Rererence code for bit packed mod6 sieve

### Pre-requisites
- [golang](https://go.dev/dl/) 
- make

### How to build
`make all`

### How to run
`./prime` to get results in a file named `bits`

```
$ ./prime
$ ls -al bits
-rw-r--r-- 1 pgp pgp 4 Feb 16 23:44 bits
$ xxd bits
00000000: fede b769
```

### How to decipher the result
```
$ xxd -b bits
00000000: 11111110 11011110 10110111 01101001

1's represent prime numbers following a mod-6 wheel starting at 5

1. 11111110 is `5 7 11 13 17 19 23` 25 is a multiple of 5 and is marked by 0.
2. 11011110 is `29 31 37 41 43 47` 35 and 49 are marked by 0
3. 10110111 is `53 59 61 67 71 73` 55 and 65 are marked by 0
4. 01101001 is `79 83 89 97`       77, 85, 91 and 95 are marked by 0 
```

### How to run a count
`make all` will build a utility called `utils`. If not present, run `make all` or `make utils`

This utility will count primes in the `bits` file if present. Default command is `count` and default file is `bits`

```
$ ./utils 
23

$ ./utils count
23

$ ./utils count bits
23

$ ./utils count bits1
open bits1: no such file or directory
```

### How to print all primes
```
$ ./utils print
5 7 11 13 17 19 23 
29 31 37 41 43 47 
53 59 61 67 71 73 
79 83 89 97
```

### Get a minimal print
Prints the byte and the starting prime for the byte
```
$ ./utils print_min
FE 5
DE 29
B7 53
69 79
```

### Other utilities
`check_times.sh` is provided to run and measure times from 10^1 to 10^9

`shasums` file contains current shasums for comparison
