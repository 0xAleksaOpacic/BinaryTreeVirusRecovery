package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// BuildTree builds a binary tree from a nested slice of values
// Example input: [1, [2, nil, nil], [3, [4, nil, nil], [5, nil, nil]]]
func BuildTree(data []interface{}) *TreeNode {
	if len(data) == 0 || data[0] == nil {
		return nil
	}

	root := &TreeNode{Val: data[0].(int)}
	if len(data) > 1 {
		if left, ok := data[1].([]interface{}); ok {
			root.Left = BuildTree(left)
		}
		if len(data) > 2 {
			if right, ok := data[2].([]interface{}); ok {
				root.Right = BuildTree(right)
			}
		}
	}
	return root
}

func TestMinRecoveryCovers_SingleNode(t *testing.T) {
	t.Parallel()

	// Construct a binary tree:
	//     1
	root := BuildTree([]interface{}{1})

	require.Equal(t, 1, minRecoveryCovers(root))
}

func TestMinRecoveryCovers_TwoNodes(t *testing.T) {
	t.Parallel()

	// Construct a binary tree:
	//     1
	//    /
	//   2
	root := BuildTree([]interface{}{1, []interface{}{2}})

	require.Equal(t, 1, minRecoveryCovers(root))
}

func TestMinRecoveryCovers_BalancedTree(t *testing.T) {
	t.Parallel()

	// Construct a binary tree:
	//     1
	//    / \
	//   2   3
	root := BuildTree([]interface{}{1, []interface{}{2}, []interface{}{3}})

	require.Equal(t, 1, minRecoveryCovers(root))
}

func TestMinRecoveryCovers_ComplexTree(t *testing.T) {
	t.Parallel()

	// Construct a binary tree:
	//       1
	//      / \
	//     2   3
	//      \ /
	//      4 5
	//        /
	//       6
	root := BuildTree([]interface{}{
		1,
		[]interface{}{2, nil, []interface{}{4}},
		[]interface{}{3, []interface{}{5, []interface{}{6}}},
	})
	require.Equal(t, 2, minRecoveryCovers(root))
}
