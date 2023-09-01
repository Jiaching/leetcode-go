/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	mapValues := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		value := nums[i]

		if _, ok := mapValues[value]; ok {
			return true
		}

		mapValues[value] = 0
	}
	return false
}

// Time : O(n)
// Space :  O(n)

// @lc code=end

