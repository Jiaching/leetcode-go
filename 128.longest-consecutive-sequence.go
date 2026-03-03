/*
 * @lc app=leetcode id=128 lang=golang
 *
 * Time: O(n)
 * Space: O(n)
*/

// @lc code=start
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	longestLength := 0
	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		node := nums[i]
		if _, ok := m[node]; ok {
			continue
		}

		leftLength := m[node-1]
		rightLength := m[node+1]
		length := leftLength + rightLength + 1
		m[node] = length
		m[node-leftLength] = length
		m[node+rightLength] = length
		if length > longestLength {
			longestLength = length
		}
	}

	return longestLength
}

// @lc code=end
