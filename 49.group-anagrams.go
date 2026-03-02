/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	mapValue := make(map[[26]byte][]string, len(strs))
	for _, str := range strs {
		letterArray := [26]byte{}
		for i := 0; i < len(str); i++ {
			letterArray[str[i]-'a']++
		}

		mapValue[letterArray] = append(mapValue[letterArray], str)
	}

	result := make([][]string, len(mapValue))
	i := 0
	for _, v := range mapValue {
		result[i] = v
		i++
	}
	return result
}

// @lc code=end

