package chunkarray

// Chunk an array into many subarrays based on the chunk size
// each subarray length equals the chunk size
// if there is not enough sub array elements to fulfill the chunk size
// the subarray will contain the maximum it can
func ChunkArray(values []int, size int) [][]int {
	var chunked [][]int
	// we init empty chunk and subarray
	var subarray []int

	for i := 0; i < len(values); i++ {
		// add every element into the subarray
		subarray = append(subarray, values[i])

		hasMaxCapacity := len(subarray) == size
		isLastIteration := i+1 == len(values)
		// if the subarray is at chunk max size
		// or this is the last iteration
		if hasMaxCapacity || isLastIteration {
			// we add the subarray to the chunked array
			chunked = append(chunked, subarray)
			// and reset the subarray
			subarray = nil
		}
	}

	return chunked
}
