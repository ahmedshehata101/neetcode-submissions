func isPalindrome(s string) bool {

	palindrom := true
	start := 0
	end := len(s) - 1
	s = strings.ToLower(s)
	for start < end {
		if (s[start] >= 'A' && s[start] <= 'Z') || (s[start] >= 'a' && s[start] <= 'z') || (s[start] >= '0' && s[start] <= '9') {
			if (s[end] >= 'A' && s[end] <= 'Z') || (s[end] >= 'a' && s[end] <= 'z') || (s[end] >= '0' && s[end] <= '9') {

				if s[start] == s[end] {
					palindrom = true
					start++
					end--
				} else {
					return false
				}
			} else {
				//fmt.Println(s[end])
				end--
				continue
			}
		} else {
			//fmt.Println(s[start])
			start++
			continue
		}
	}
	return palindrom
}
