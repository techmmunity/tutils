package tutils

import (
	"math"
)

func Chunk[T any](arr []T, itemsPerChunk int) [][]T {
	arrLen := len(arr)
	arrLastItem := arrLen
	amountOfChunks := int(math.Ceil(float64(arrLen) / float64(itemsPerChunk)))

	var chunks [][]T

	for i := 0; i < amountOfChunks; i++ {
		startIdx := i * itemsPerChunk;
		endIdx := startIdx + itemsPerChunk;

		if (endIdx > arrLastItem) {
			endIdx = arrLastItem
		}

		chunks = append(chunks, arr[startIdx:endIdx])
	}

	return chunks
}
