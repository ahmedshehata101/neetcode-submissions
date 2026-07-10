func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)
	var freq [26]int
	for _, s := range strs {

		for _, ch := range s {
			freq[ch-'a']++
		}
		groups[freq] = append(groups[freq], s)
		freq = [26]int{}
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}