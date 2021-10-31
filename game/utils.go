package game

func min(nums ...int) (int, int) {
	idx, minval := 0, nums[0]
	for i, val := range nums {
		if val < minval {
			idx = i
			minval = val
		}
	}
	return idx, minval
}

func max(nums ...int) (int, int) {
	idx, maxval := 0, nums[0]
	for i, val := range nums {
		if val > maxval {
			idx = i
			maxval = val
		}
	}
	return idx, maxval
}
