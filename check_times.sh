#!/bin/sh
pgm=${1:-./c/prime/prime}
max=${2:-9}

echo "$pgm"

for i in $(seq 2 "$max"); do
  j=$(echo "10^$i" | bc)
  echo 10^"$i"
  /usr/bin/time -f "\t%E real,\t%U user,\t%S sys" "$pgm" "$j"
done
