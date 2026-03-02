/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	mapValue := make(map[[26]int][]string, len(strs))
	for _, str := range strs {
		letterArray := [26]int{}
		for _, v := range str {
			letterArray[v-'a']++
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

