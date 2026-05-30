package algorithmes

func MergeSort(data []int) []int {

	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2

	left := MergeSort(data[:mid])
	right := MergeSort(data[mid:])

	return merge(left, right)
}

func merge(l []int, r []int) []int {

	result := make([]int, 0)

	i, j := 0, 0

	for i < len(l) && j < len(r) {

		if l[i] < r[j] {
			result = append(result, l[i])
			i++
		} else {
			result = append(result, r[j])
			j++
		}
	}

	result = append(result, l[i:]...)
	result = append(result, r[j:]...)

	return result
}
