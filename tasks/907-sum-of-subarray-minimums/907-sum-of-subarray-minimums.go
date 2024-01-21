package _907_sum_of_subarray_minimums

func sumSubarrayMins(arr []int) int {
	sum := 0

	mins := make([]int, len(arr))
	copy(mins, arr)

	for _, el := range arr {
		sum += el
	}

	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j+i] < mins[j] {
				mins[j] = arr[j+i]
			}

			sum += mins[j]
		}
	}

	return sum % 1000000007
}
