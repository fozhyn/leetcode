/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode-cn.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (46.83%)
 * Likes:    314
 * Dislikes: 0
 * Total Accepted:    192.3K
 * Total Submissions: 410.7K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 *
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 *
 * 示例 1:
 *
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "race a car"
 * 输出: false
 *
 *
 */

// @lc code=start

func isNumAndLetter(b byte) bool {
	return b >= '0' && b <= '9' || b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}

func isPalindrome(s string) bool {
	length := len(s)
	if length == 0 {
		return true
	}

	i, j := 0, length-1
	for i < j {
		if !isNumAndLetter(s[i]) {
			i++
			continue
		}

		if !isNumAndLetter(s[j]) {
			j--
			continue
		}

		if !(s[i] == s[j] || (s[i] >= 'A' && s[j] >= 'A' && (s[i]-s[j] == 'a'-'A' || s[j]-s[i] == 'a'-'A'))) {
			break
		}
		i++
		j--
	}

	return i >= j
}

// @lc code=end

