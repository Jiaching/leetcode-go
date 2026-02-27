/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, v := range nums {

		if value, exists := numMap[v]; exists {
			return []int{value, index}
		}

		complement := target - v
		numMap[complement] = index
	}
	return nil
}

// @lc code=end

