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

	mapValue := make(map[string]int, len(t))
	str := strings.Split(t, "")
	// Put character of string into map
	for _, str := range str {
		mapValue[str] += 1
	}

	foundStr := strings.Split(s, "")
	for _, target := range foundStr {
		if _, ok := mapValue[target]; ok {
			mapValue[target] -= 1
			if mapValue[target] == 0 {
				delete(mapValue, target)
			}
		}
	}

	// for key, element := range mapValue {
	// 	fmt.Println("Key:", key, "=>", "Element:", element)
	// }

	return len(mapValue) == 0
	// Remove items if exists.

	// Return ture if mapValue is empty/
}

// @lc code=end

