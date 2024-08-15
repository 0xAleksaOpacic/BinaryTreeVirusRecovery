package main

import (
	"fmt"
)

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// minRecoveryCovers returns the minimum number of computers needed to run the recovery software.
func minRecoveryCovers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := []*TreeNode{}
	curr := root
	var lastVisited *TreeNode
	installCount := 0

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			node := stack[len(stack)-1]

			if node.Right != nil && lastVisited != node.Right {
				curr = node.Right
			} else {
				stack = stack[:len(stack)-1]
				lastVisited = node

				// Handle the antivirus placement logic
				if handleAntivirusPlacement(node) == 2 {
					installCount++
					logAntivirusPlacement(node.Val)
				}
			}
		}
	}

	// If the root itself is not covered, place antivirus
	if getState(root) == 0 {
		logAntivirusPlacement(root.Val)
		installCount++
	}

	return installCount
}

// handleAntivirusPlacement checks the status of the children and decides whether to place antivirus software
func handleAntivirusPlacement(node *TreeNode) int {
	leftState := getState(node.Left)
	rightState := getState(node.Right)

	if leftState == 0 || rightState == 0 {
		// If either child is not covered, place recovery software here
		setNodeState(node, 2)
		return 2
	} else if leftState == 2 || rightState == 2 {
		// If either child has antivirus software, this node is covered
		setNodeState(node, 1)
		return 1
	} else {
		// Node is not covered
		setNodeState(node, 0)
		return 0
	}
}

// logAntivirusPlacement logs the placement of antivirus software
func logAntivirusPlacement(nodeVal int) {
	fmt.Println("Antivirus set on computer with id:", nodeVal)
}

// Global map to store the state of the nodes
var nodeState = make(map[*TreeNode]int)

// getState retrieves the state of the node
func getState(node *TreeNode) int {
	if node == nil {
		return 1 // Null nodes are considered covered
	}
	if state, exists := nodeState[node]; exists {
		return state
	}
	return 0 // Default state is not covered
}

// setNodeState sets the state of the node
func setNodeState(node *TreeNode, state int) {
	if node != nil {
		nodeState[node] = state
	}
}

func main() {
	// Example usage:
	// Construct a binary tree:
	//     1
	//    / \
	//   2   3
	//  / \   \
	// 4   5   6

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	// Calculate the minimum recovery covers
	result := minRecoveryCovers(root)
	fmt.Println("Minimum number of computers needed:", result)
}
