#!/bin/sh
j=1
for i in $(seq 9); do
  j=$((j*10))
  echo 10^"$i"
  /usr/bin/time -f "\t%E real,\t%U user,\t%S sys" ./prime $j
  shasum bits
done

