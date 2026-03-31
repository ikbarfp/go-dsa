package sorting

// BubbleSort
//
//	Time complexity 	= O(n^2)
//	Space complexity 	=
func BubbleSort(array []int) []int {
	if array == nil || len(array) <= 1 {
		return array
	}

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}
