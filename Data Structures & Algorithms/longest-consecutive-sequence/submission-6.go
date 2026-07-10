import "slices"
func longestConsecutive(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	if len(nums) == 0 {
		return 0
	}
	slices.Sort(nums)
	counter := 1
	streak := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i] {
			continue
		}
		if nums[i+1]-nums[i] == 1 {
			counter++
		} else {
			counter = 1
		}
		if streak < counter {
			streak = counter
		}
	}
	return streak
}