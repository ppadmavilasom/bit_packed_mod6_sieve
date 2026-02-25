#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <errno.h>
#include <math.h>
#include <string.h>
#include "prototypes.h"

const int WHEEL_SIZE     = 6;
const int SQUARE_TOGGLE  = 2;
const int UNITS_PER_BYTE = 24;
const int DEFAULT_LIMIT  = 100;

static const int BIT_SHIFTS[8] = {
	1<<7, 1<<6, 1<<5, 1<<4,
	1<<3, 1<<2, 1<<1, 1<<0
};

static const int BIT_MASKS[8] = {
	~(1<<7), ~(1<<6), ~(1<<5), ~(1<<4),
	~(1<<3), ~(1<<2), ~(1<<1), ~(1<<0)
};

int main(int argc, char **argv) {
	uint64_t limit = 0;
	char *pBytes = NULL;
	size_t len = 0;
	int result = getArg(1, argc, argv, DEFAULT_LIMIT, &limit);
	if (result > 0) {
		goto error;
	}

	result = calcPrimes(limit, &pBytes, &len);
	if (result > 0) {
		goto error;
	}
	result = writeFile("bits", pBytes, len);
	if (result > 0) {
		goto error;
	}

cleanup:
	return result;

error:
	if (pBytes != NULL) {
		free(pBytes);
		pBytes = NULL;
	}
	goto cleanup;
}

int calcPrimes(uint64_t limit, char **ppBytes, size_t *pLen) {
	int result = 0;
	size_t len = 0;
	char *pBytes = NULL;
	OFFSETS o = {5, 7, 3, 7^3, 4, 0, 0, 1, 1, 2, 1};

	if (limit < UNITS_PER_BYTE || ppBytes == NULL || pLen == NULL) {
		result = EINVAL;
		goto error;
	}

	len = limit / UNITS_PER_BYTE;
	pBytes = malloc(len);
	if (pBytes == NULL) {
		result = EINVAL;
		goto error;
	}

	memset(pBytes, 0xFF, len);

	uint64_t sqrtN = (uint64_t)sqrt(limit);
	for(; o.start <= sqrtN; next(&o)) {
		if ((pBytes[o.index >> 3] & BIT_SHIFTS[o.index & 7]) == 0) {
			continue;
		}
		uint64_t offset = o.offset1;
		uint64_t bit = ((o.squareIndex + 1) << 3) - 1;
		uint64_t row = bit >> 3;
		if ((o.index & 1) == 1) {
			offset = o.offset2;
		}

		while (row < len) {
			pBytes[row] &= BIT_MASKS[bit & 7];
			offset ^= o.offsetSwap;
			bit += offset;
			row = bit >> 3;
		}
	}

	*ppBytes = pBytes;
	*pLen = len;
cleanup:
	return result;

error:
	if (pBytes != NULL) {
		free(pBytes);
	}
	if (ppBytes != NULL) {
		*ppBytes = NULL;
	}
	if (pLen != NULL) {
		*pLen = 0;
	}
	goto cleanup;
}

void next(POFFSETS pO) {
	pO->inc ^= WHEEL_SIZE;
	pO->start += pO->inc;
	pO->offset2 += 2;
	pO->offset1 = (pO->start << 1) - pO->offset2;
	pO->offsetSwap = pO->offset1 ^ pO->offset2;
	pO->index++;

	// Square optimization updates
	pO->squareIndex += pO->squareInc;
	if (pO->squareSeq == 2) {
		pO->squareSeq2 += 2;
		pO->squareInc = pO->squareSeq2;
	} else {
		pO->squareSeq1++;
		pO->squareInc = pO->squareSeq1;
	}
	pO->squareSeq ^= SQUARE_TOGGLE;
}

int writeFile(const char *pName, const char *pBytes, const size_t len) {
	int result = 0;
	size_t bytesWritten = 0;
	if (pName == NULL || pBytes == NULL) {
		result = EINVAL;
		goto error;
	}
	FILE *fp = fopen(pName, "wb");
	if (fp == NULL) {
		result = ENOENT;
		goto error;
	}
	bytesWritten = fwrite(pBytes, 1, len, fp);
	if (bytesWritten != len) {
		result = EINVAL;
		goto error;
	}
cleanup:
	return result;

error:
	if (fp != NULL) {
		fclose(fp);
		fp = NULL;
	}
	goto cleanup;
}

int getArg(int index, int argc, char **argv, uint64_t maxLimit, uint64_t *pLimit) {
	int result = 0;
	uint64_t limit = 0;
	if (pLimit == NULL || argv == NULL) {
		result = EINVAL;
		goto error;
	}
	if (argc > index) {
		limit = atoll(argv[index]);
	}
	if (limit < maxLimit) {
		limit = maxLimit;
	}

	*pLimit = limit;
cleanup:
	return result;

error:
	goto cleanup;
}
