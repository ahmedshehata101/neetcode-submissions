func hasDuplicate(nums []int) bool {
	haveseen := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if _, ok := haveseen[nums[i]]; ok {
			return true
		} else {
			haveseen[nums[i]] = false
		}
	}
	return false
}
