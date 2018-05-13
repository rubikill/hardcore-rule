package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var charArr = make([]int, 26)
	for i := 0; i < len(s); i++ {
		charArr[s[i]-97]++
		charArr[t[i]-97]--
	}

	for i := 0; i < len(charArr); i++ {
		if charArr[i] > 0 {
			return false
		}
	}

	return true
}
