package easy

import "container/list"

//给定一个二叉树的根节点 root ，返回它的 中序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[2,1]
//
//
// 示例 5：
//
//
//输入：root = [1,null,2]
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1134 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	inorder(root, &ans)
	return ans
}

func inorder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, ans)
	*ans = append(*ans, root.Val)
	inorder(root.Right, ans)
}

func iterInorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}

	nodeStack := list.New()
	for root != nil {
		nodeStack.PushBack(root)
		root = root.Left
	}

	for nodeStack.Len() > 0 {
		nodeElement := nodeStack.Back()
		//if _, ok:=curNode.Value.(*TreeNode);!ok{
		//	return err
		//}
		curNode := nodeElement.Value.(*TreeNode)
		ans = append(ans, curNode.Val)
		nodeStack.Remove(nodeElement)

		if curNode.Right != nil {
			root := curNode.Right
			for root != nil {
				nodeStack.PushBack(root)
				root = root.Left
			}
		}

	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
