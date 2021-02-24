/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (73.53%)
 * Likes:    326
 * Dislikes: 0
 * Total Accepted:    158.3K
 * Total Submissions: 215.2K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出：[2]
 *
 *
 * 示例 2：
 *
 * 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出：[9,4]
 *
 *
 *
 * 说明：
 *
 *
 * 输出结果中的每个元素一定是唯一的。
 * 我们可以不考虑输出结果的顺序。
 *
 *
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	mapping := make(map[int]bool)
	result := make([]int, 0)
	for _, v := range nums1 {
		mapping[v] = true
	}

	for _, v := range nums2 {
		if mapping[v] {
			result = append(result, v)
			mapping[v] = false
		}
	}

	return result
}

// @lc code=end

