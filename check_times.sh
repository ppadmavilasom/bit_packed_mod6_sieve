#!/bin/sh
pgm=${1:-./prime/prime}
for i in $(seq 2 9); do
  j=$(echo "10^$i" | bc)
  echo 10^"$i"
  /usr/bin/time -f "\t%E real,\t%U user,\t%S sys" "$pgm" "$j"
done

