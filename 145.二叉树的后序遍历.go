// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)
	return append(append(left, right...), root.Val)
}*/

func postorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	var lastVisited *TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisited {
			res = append(res, node.Val)
			stack = stack[:len(stack)-1]
			lastVisited = node
			node = nil
		} else {
			node = node.Right
		}
	}
	return res
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode.cn/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Easy (76.57%)
 * Likes:    1173
 * Dislikes: 0
 * Total Accepted:    771.9K
 * Total Submissions: 1M
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,3]
 * 输出：[3,2,1]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [0, 100] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶：递归算法很简单，你可以通过迭代算法完成吗？
 *
 */


