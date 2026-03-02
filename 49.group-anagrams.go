/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	mapValue := make(map[string][]string, len(strs))
	for _, str := range strs {
		key := stringEncode(str)
		mapValue[key] = append(mapValue[key], str)
	}

	result := make([][]string, len(mapValue))
	i := 0
	for _, v := range mapValue {
		result[i] = v
		i++
	}
	return result
}

func stringEncode(str string) string {
	if str == "" {
		return ""
	}

	letterArray := [26]int{}
	for _, v := range str {
		letterArray[v-'a']++
	}

	var sb strings.Builder
	for i := 0; i < len(letterArray); i++ {
		if letterArray[i] == 0 {
			continue
		}
		sb.WriteString(fmt.Sprintf("%s%d", string(i+'a'), letterArray[i]))
	}

	return sb.String()

}

// @lc code=end

