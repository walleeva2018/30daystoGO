### Problem 

Given the root of a binary tree, return the inorder traversal of its nodes' values.

[link](https://leetcode.com/problems/binary-tree-inorder-traversal/)

#### what is an inorder?

Inorder traversal is one of the methods used to traverse (visit all the nodes of) a binary tree. The key characteristic of this traversal method is that it visits nodes in a specific order: Left → Root → Right.



       4
      / \
     2   6
    / \  / \
   1  3 5   7

Performing an inorder traversal on this tree will give:

- Start with the left subtree of 4 → Traverse 2.
- Within the subtree of 2, traverse 1 (leftmost node).
- Visit the root 2, then traverse its right subtree 3.
- Visit the root 4.
- Move to the right subtree of 4 → Traverse 6.
- Within the subtree of 6, traverse 5 (left child), visit 6, then traverse 7.

Code 

```
func inorderTraversal(root *TreeNode) []int {
    var result []int
    
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
```

- Creates an empty result slice
- Defines an inner recursive traverse function
- Handles nil nodes by returning
- Recursively traverses left subtree
- Adds current node's value to result
- Recursively traverses right subtree
- Calls traverse with root node
- Returns the complete inorder traversal res