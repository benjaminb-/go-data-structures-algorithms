package chunkarray

import (
	"reflect"
	"testing"
)

type chunkConfig struct {
	array []int
	size  int
}

func TestChunkArray(t *testing.T) {

	chunkArrayTests := []struct {
		test string
		got  chunkConfig
		want [][]int
	}{{test: "chunk in 2 subarrays with max size of 3", got: chunkConfig{
		array: []int{1, 2, 3, 4, 5, 6}, size: 3},
		want: [][]int{{1, 2, 3}, {4, 5, 6}}},
		{test: "chunk in 3 subarrays with max size of 2", got: chunkConfig{
			array: []int{1, 2, 3, 4, 5, 6}, size: 2},
			want: [][]int{{1, 2}, {3, 4}, {5, 6}}},
		{test: "chunk in 1 subarrays with max size of 10", got: chunkConfig{
			array: []int{1, 2, 3, 4, 5, 6}, size: 10},
			want: [][]int{{1, 2, 3, 4, 5, 6}}},
		{test: "chunk in 2 subarrays with max size of 5", got: chunkConfig{
			array: []int{1, 2, 3, 4, 5, 6}, size: 5},
			want: [][]int{{1, 2, 3, 4, 5}, {6}}},
	}

	for _, tt := range chunkArrayTests {
		t.Run(tt.test, func(t *testing.T) {
			if reflect.DeepEqual(ChunkArray(tt.got.array, tt.got.size), tt.want) == false {
				t.Error("failed ChunkArray", ChunkArray(tt.got.array, tt.got.size), tt.want)
			}
		})
	}

}
