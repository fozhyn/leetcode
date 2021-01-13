/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (43.55%)
 * Likes:    2086
 * Dislikes: 0
 * Total Accepted:    500.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */

// @lc code=start
func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, v := range s {
		// 左括号直接入栈
		if v == '[' || v == '(' || v == '{' {
			stack = append(stack, v)
			continue
		}

		// 遇到右括号，如果stack长度为0，则肯定无效
		if len(stack) == 0 {
			return false
		}

		// 括号不成对，直接返回无效
		last := stack[len(stack)-1]
		if !(v == ')' && last == '(' ||
			v == ']' && last == '[' ||
			v == '}' && last == '{') {
			return false
		}

		//遇到成对括号，将左括号从栈中移除
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

// @lc code=end

