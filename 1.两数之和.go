/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	for i, v := range nums {
		if idx, ok := mapping[target-v]; ok {
			return []int{idx, i}
		}
		mapping[v] = i
	}

	return nil

}

// @lc code=end

