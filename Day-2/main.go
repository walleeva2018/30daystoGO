package main

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    result := []int{}
    
    var traverse func(*TreeNode)
    traverse = func(node *TreeNode) {
        if node == nil {
            return
        }
        traverse(node.Left)
        result = append(result, node.Val)
        traverse(node.Right)
    }
    
    traverse(root)
    return result
}

func main() {
   
    root := &TreeNode{Val: 4}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 6}
    root.Left.Left = &TreeNode{Val: 1}
    root.Left.Right = &TreeNode{Val: 3}
    root.Right.Left = &TreeNode{Val: 5}
    root.Right.Right = &TreeNode{Val: 7}

    
    result := inorderTraversal(root)
    
    
    fmt.Println(result)
}