func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		if _, ok := freq[n]; ok {
			freq[n]++
		} else {
			freq[n] = 1
		}
	}
	buckets := make([][]int, len(nums)+1)
	for k, v := range freq {
		buckets[v] = append(buckets[v], k)
	}
	results := []int{}
	for i := len(buckets) - 1; i >= 0 && len(results) < k; i-- {
		for _, num := range buckets[i] {
			results = append(results, num)
		}
		if len(results) == k {
			break
		}
	}
	return results
}