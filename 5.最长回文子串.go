/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (33.15%)
 * Likes:    3240
 * Dislikes: 0
 * Total Accepted:    485.2K
 * Total Submissions: 1.5M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出："bb"
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "a"
 * 输出："a"
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "ac"
 * 输出："a"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 仅由数字和英文字母（大写和/或小写）组成
 *
 *
 */

// @lc code=start
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	ans := ""
	for i := 0; i < n; i++ {
		for j := 0; j+i < n; j++ {
			if i == 0 {
				// 长度为1的字符串是回文串
				dp[j][j+i] = 1
			} else if i == 1 {
				// 长度为2的字符串，如果两个字符相等，是回文串
				if s[j] == s[j+i] {
					dp[j][j+i] = 1
				} else {
					dp[j][j+i] = 0
				}
			} else {
				// 长度超过2的字符串，如果首尾字符相等
				// 是否是回文串取决于掐头去尾之后的子串是否是回文串
				if s[j] == s[j+i] {
					dp[j][j+i] = dp[j+1][j+i-1]
				} else {
					dp[j][j+i] = 0
				}
			}

			if dp[j][j+i] > 0 && i+1 > len(ans) {
				ans = s[j : j+i+1]
			}
		}
	}

	return ans
}

// @lc code=end

