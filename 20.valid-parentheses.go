/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	stack := list.New()
	strs := strings.Split(s, "")

	if len(s)%2 == 1 {
		return false
	}

	for _, str := range strs {
		if stack.Len() == 0 {
			stack.PushBack(str)
			continue
		}

		peekStr := stack.Back()
		if str == "]" && peekStr.Value == "[" {
			stack.Remove(peekStr)
		} else if str == ")" && peekStr.Value == "(" {
			stack.Remove(peekStr)
		} else if str == "}" && peekStr.Value == "{" {
			stack.Remove(peekStr)
		} else {
			stack.PushBack(str)
		}
	}

	return stack.Len() == 0

}

// @lc code=end
