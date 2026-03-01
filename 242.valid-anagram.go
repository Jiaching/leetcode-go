/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	cMap := make(map[rune]int, 26)

	for _, c := range s {
		cMap[c]++
	}

	for _, c := range t {
		cMap[c]--
		if cMap[c] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end

