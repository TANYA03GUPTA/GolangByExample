package main

import("fmt")

type TreeNode struct{
	val int
	left *TreeNode
	right *TreeNode
}
func pre(root *TreeNode)[]int{
	if root == nil {
		return nil
	}
	l := pre(root.left)
	r := pre(root.right)

	output := make([]int,0)
	output = append(output, root.val)
	output = append(output, l...)
	output = append(output, r...)
	return output
}

func main(){
	root := TreeNode{val:1}
	root.left = &TreeNode{val:2}
	root.right = &TreeNode{val:3}
	root.left.left = &TreeNode{val :5}
	root.left.right = &TreeNode{val:6}
	output := pre(&root)
	fmt.Printf("Preorder Traversal:",output)
}