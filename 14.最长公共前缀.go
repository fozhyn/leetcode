/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (39.07%)
 * Likes:    1414
 * Dislikes: 0
 * Total Accepted:    427.6K
 * Total Submissions: 1.1M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：strs = ["flower","flow","flight"]
 * 输出："fl"
 *
 *
 * 示例 2：
 *
 *
 * 输入：strs = ["dog","racecar","car"]
 * 输出：""
 * 解释：输入不存在公共前缀。
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * 0
 * strs[i] 仅由小写英文字母组成
 *
 *
 */

// @lc code=start

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}

	return str1[:index]
}

func longestCommonPrefix(strs []string) string {
	counts := len(strs)
	if counts == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < counts; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}

	return prefix
}

// @lc code=end

