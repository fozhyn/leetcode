/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (39.64%)
 * Likes:    3730
 * Dislikes: 0
 * Total Accepted:    338.4K
 * Total Submissions: 853.8K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
 *
 * 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums1 = [1,3], nums2 = [2]
 * 输出：2.00000
 * 解释：合并数组 = [1,2,3] ，中位数 2
 *
 *
 * 示例 2：
 *
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：2.50000
 * 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 *
 *
 * 示例 3：
 *
 * 输入：nums1 = [0,0], nums2 = [0,0]
 * 输出：0.00000
 *
 *
 * 示例 4：
 *
 * 输入：nums1 = [], nums2 = [1]
 * 输出：1.00000
 *
 *
 * 示例 5：
 *
 * 输入：nums1 = [2], nums2 = []
 * 输出：2.00000
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 *
 *
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i := findKth(nums1, nums2, (len(nums1)+len(nums2)+1)/2)
	j := findKth(nums1, nums2, (len(nums1)+len(nums2)+2)/2)

	return float64(i+j) / float64(2.0)
}
func min(i, j int) int {
	if i > j {
		return j
	}

	return i
}

func findKth(nums1 []int, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}

	if len(nums2) == 0 {
		return nums1[k-1]
	}

	if k == 1 {
		return min(nums1[0], nums2[0])
	}

	m := nums1[min(k/2, len(nums1))-1]
	n := nums2[min(k/2, len(nums2))-1]

	if m < n {
		return findKth(nums1[min(k/2, len(nums1)):], nums2, k-min(k/2, len(nums1)))
	}

	return findKth(nums1, nums2[min(k/2, len(nums2)):], k-min(k/2, len(nums2)))
}

// @lc code=end

