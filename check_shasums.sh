#!/bin/sh
# shellcheck disable=SC2013
pgm=${1:-./prime/prime}
max=${2:-9}

index=1
power=2
limit=100
for shasum in $(cut -d' ' -f1 shasums); do
  if [ "$index" -gt "$max" ]; then
    break
  fi
  current_shasum=$("$pgm" "$limit" && shasum bits | cut -d' ' -f1)
  if [ "$current_shasum" != "$shasum" ]; then
    echo "limit=10^$power. Expected: $shasum. Got: $current_shasum"
    echo "Check Failed."
    exit 1
  fi
  echo "limit=10^$power. $current_shasum match."
  power=$((power+1))
  limit=$((limit*10))
  index=$((index+1))
done
echo "Check Succeeded."
echo
