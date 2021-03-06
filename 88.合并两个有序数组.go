/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 *
 * https://leetcode-cn.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (49.12%)
 * Likes:    736
 * Dislikes: 0
 * Total Accepted:    254.2K
 * Total Submissions: 517.4K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
 *
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 有足够的空间（空间大小等于 m + n）来保存 nums2
 * 中的元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
 * 输出：[1,2,2,3,5,6]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [1], m = 1, nums2 = [], n = 0
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * 1
 * nums1.length == m + n
 * nums2.length == n
 * -10^9
 *
 *
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, p := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[p] = nums1[i]
			i--
		} else {
			nums1[p] = nums2[j]
			j--
		}
		p--
	}

	for j >= 0 {
		nums1[p] = nums2[j]
		p--
		j--
	}
}

// @lc code=end

