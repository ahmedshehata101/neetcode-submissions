func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	output := make([]int, len(nums))
	prefix[0] = 1
	suffix[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {

		prefix[i] = prefix[i-1] * nums[i-1]

	}
	for x := len(nums) - 2; x >= 0; x-- {
		suffix[x] = suffix[x+1] * nums[x+1]
	}
	for i := 0; i < len(nums); i++ {
		output[i] = prefix[i] * suffix[i]
	}
	return output
}