/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	mapValues := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		wantedValue := target - nums[i]
		if position, ok := mapValues[wantedValue]; ok {
			return []int{position, i}
		}
		mapValues[nums[i]] = i
	}

	return nil
}

// @lc code=end

