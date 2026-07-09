func search(nums []int, target int) int {
    	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1   // move left pointer
		} else {
			right = mid - 1  // move right pointer
		}
	}

	return -1
}
