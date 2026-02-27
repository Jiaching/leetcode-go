/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	valueMap := make(map[int]interface{})

	for _, v := range nums {
		if _, exists := valueMap[v]; exists {
			return true
		}
		valueMap[v] = nil
	}
	return false
}

// Time : O(n)
// Space :  O(n)

// @lc code=end

