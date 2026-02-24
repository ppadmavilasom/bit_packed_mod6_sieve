#pragma once

typedef struct Offsets {
	uint64_t start;
	uint64_t offset1;
	uint64_t offset2;
	uint64_t offsetSwap;
	uint64_t inc;
	uint64_t index;
} OFFSETS, *POFFSETS;

void next(POFFSETS pO);
int writeFile(const char *pName, const char *pBytes, const size_t len);
int calcPrimes(uint64_t limit, char **ppBytes, size_t *pLen);
int getArg(int index, int argc, char **argv, uint64_t maxLimit, uint64_t *limit);
