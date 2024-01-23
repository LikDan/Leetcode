package _1_two_sum

func twoSum(nums []int, target int) []int {
	d := map[int]int{}

	for i, num := range nums {
		if d[target-num] != 0 {
			return []int{d[target-num] - 1, i}
		}

		d[num] = i + 1
	}

	return nil
}
