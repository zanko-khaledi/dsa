package algorithmes

func BinarySearch(data []int, target int) int {

	left, right := 0, len(data)-1

	for left <= right {
		mid := left + (right-left)/2

		if data[mid] == target {
			return mid
		} else if data[mid] < target {
			left++
		} else {
			right--
		}
	}

	return -1
}
