/*
二叉树：二叉树是一种树形结构，每个节点最多只有两个子节点。
主要思想：
	1. 创建左右两个指针节点指针。
	2. 使用左右节点指针完成前、中、后序遍历、搜索,删除。
	3. 前序：父节点=>左子节点=>右子节点。
	4. 中序：左子节点=>父节点=>右子节点。
	5. 后序：左子节点=>右子节点=>父节点。
*/
package tree

import "fmt"

// 二叉树节点
type BinaryTreeNode struct {
	Id int
	Data int
	Left *BinaryTreeNode
	Right *BinaryTreeNode
}

// 删除节点
func (binaryTreeNode *BinaryTreeNode) DeleteNode(id int)  {
	if binaryTreeNode.Left!=nil && binaryTreeNode.Left.Id == id{
		binaryTreeNode.Left = nil
		return
	}
	if binaryTreeNode.Right!=nil && binaryTreeNode.Right.Id == id{
		binaryTreeNode.Right = nil
		return
	}
	if binaryTreeNode.Left!=nil{
		binaryTreeNode.Left.DeleteNode(id)
	}
	if binaryTreeNode.Right!=nil{
		binaryTreeNode.Right.DeleteNode(id)
	}
}

// 前序遍历
func (binaryTreeNode *BinaryTreeNode) PreorderTraversal()  {
	fmt.Print(binaryTreeNode.Id,"=>")
	if binaryTreeNode.Left != nil{
		binaryTreeNode.Left.PreorderTraversal()
	}
	if binaryTreeNode.Right != nil{
		binaryTreeNode.Right.PreorderTraversal()
	}
}

// 中序遍历
func (binaryTreeNode *BinaryTreeNode) InorderTraversal()  {
	if binaryTreeNode.Left != nil{
		binaryTreeNode.Left.InorderTraversal()
	}
	fmt.Print(binaryTreeNode.Id,"=>")
	if binaryTreeNode.Right != nil{
		binaryTreeNode.Right.InorderTraversal()
	}
}

// 后序遍历
func (binaryTreeNode *BinaryTreeNode) PostOrderTraversal()  {
	if binaryTreeNode.Left != nil{
		binaryTreeNode.Left.PostOrderTraversal()
	}
	if binaryTreeNode.Right != nil{
		binaryTreeNode.Right.PostOrderTraversal()
	}
	fmt.Print(binaryTreeNode.Id,"=>")
}

// 前序搜索
func (binaryTreeNode *BinaryTreeNode) PreorderSearch(id int) *BinaryTreeNode {
	if binaryTreeNode.Id == id{
		return binaryTreeNode
	}
	var node *BinaryTreeNode
	if binaryTreeNode.Left != nil{
		node = binaryTreeNode.Left.PreorderSearch(id)
	}
	if node!=nil{
		return node
	}
	if binaryTreeNode.Right != nil{
		node = binaryTreeNode.Right.PreorderSearch(id)
	}
	return node
}

// 中序搜索
func (binaryTreeNode *BinaryTreeNode) InorderSearch(id int) *BinaryTreeNode {
	var node *BinaryTreeNode
	if binaryTreeNode.Left != nil{
		node = binaryTreeNode.Left.InorderSearch(id)
	}
	if node!=nil{
		return node
	}
	if binaryTreeNode.Id == id{
		return binaryTreeNode
	}
	if binaryTreeNode.Right != nil{
		node = binaryTreeNode.Right.InorderSearch(id)
	}
	return node
}

// 后序搜索
func (binaryTreeNode *BinaryTreeNode) PostOrderSearch(id int) *BinaryTreeNode {
	var node *BinaryTreeNode
	if binaryTreeNode.Left != nil{
		node = binaryTreeNode.Left.PostOrderSearch(id)
	}
	if node!=nil{
		return node
	}
	if binaryTreeNode.Right != nil{
		node = binaryTreeNode.Right.PostOrderSearch(id)
	}
	if node!=nil{
		return node
	}
	if binaryTreeNode.Id == id{
		return binaryTreeNode
	}
	return node
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

// 删除节点
func (binaryTree *BinaryTree) DeleteNode(id int)  {
	if binaryTree.Root == nil{
		fmt.Println("This tree is empty!")
	}else if binaryTree.Root.Id == id{
		binaryTree.Root = nil
	}else{
		binaryTree.Root.DeleteNode(id)
	}
}

// 前序遍历
func (binaryTree *BinaryTree) PreorderTraversal()  {
	if binaryTree.Root == nil{
		fmt.Println("This tree is empty!")
	}else{
		binaryTree.Root.PreorderTraversal()
	}
}

// 中序遍历
func (binaryTree *BinaryTree) InorderTraversal()  {
	if binaryTree.Root == nil{
		fmt.Println("This tree is empty!")
	}else{
		binaryTree.Root.InorderTraversal()
	}
}

// 后序遍历
func (binaryTree *BinaryTree) PostOrderTraversal()  {
	if binaryTree.Root == nil{
		fmt.Println("This tree is empty!")
	}else{
		binaryTree.Root.PostOrderTraversal()
	}
}

// 前序搜索
func (binaryTree *BinaryTree) PreorderSearch(id int) *BinaryTreeNode {
	if binaryTree.Root == nil{
		panic("This tree is empty!")
	}else{
		return binaryTree.Root.PreorderSearch(id)
	}
}

// 中序搜索
func (binaryTree *BinaryTree) InorderSearch(id int) *BinaryTreeNode {
	if binaryTree.Root == nil{
		panic("This tree is empty!")
	}else{
		return binaryTree.Root.InorderSearch(id)
	}
}

// 后序搜索
func (binaryTree *BinaryTree) PostOrderSearch(id int) *BinaryTreeNode {
	if binaryTree.Root == nil{
		panic("This tree is empty!")
	}else{
		return binaryTree.Root.PostOrderSearch(id)
	}
}