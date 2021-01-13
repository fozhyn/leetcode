/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 *
 * https://leetcode-cn.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (34.87%)
 * Likes:    2450
 * Dislikes: 0
 * Total Accepted:    556K
 * Total Submissions: 1.6M
 * Testcase Example:  '123'
 *
 * 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 *
 *
 *
 * 注意：
 *
 *
 * 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回
 * 0。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：x = 123
 * 输出：321
 *
 *
 * 示例 2：
 *
 *
 * 输入：x = -123
 * 输出：-321
 *
 *
 * 示例 3：
 *
 *
 * 输入：x = 120
 * 输出：21
 *
 *
 * 示例 4：
 *
 *
 * 输入：x = 0
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * -2^31
 *
 *
 */

// @lc code=start
func reverse(x int) int {
	max := 1<<31 - 1
	min := -1 << 31
	y := 0
	for x != 0 {
		if y > max/10 || (y == max/10) && (x%10) > 7 {
			return 0
		}

		if y < min/10 || (y == min/10) && (x%10) < -8 {
			return 0
		}
		y = y*10 + x%10
		x = x / 10
	}

	return y
}

// @lc code=end
