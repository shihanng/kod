package mergetree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func stringToTreeNode(input string) *TreeNode {
	builder := treeNodeBuilder{raw: input}
	return builder.build()
}

type treeNodeBuilder struct {
	raw   string
	queue []*TreeNode
}

func (b *treeNodeBuilder) nextValue() *int {
	if b.raw == "" {
		return nil
	}

	splitted := strings.SplitN(b.raw, ",", 2)

	if len(splitted) == 1 {
		b.raw = ""
	} else {
		b.raw = splitted[1]
	}

	rawValue := splitted[0]

	if rawValue == "null" || rawValue == "" {
		return nil
	}

	val, err := strconv.Atoi(rawValue)
	if err != nil {
		panic(err)
	}

	return &val
}

func (b *treeNodeBuilder) build() *TreeNode {
	val := b.nextValue()
	if val == nil {
		return nil
	}

	root := TreeNode{Val: *val}

	b.queue = append(b.queue, &root)

	for len(b.queue) > 0 {
		var node *TreeNode
		node, b.queue = b.queue[0], b.queue[1:]

		if val := b.nextValue(); val != nil {
			leftNode := TreeNode{Val: *val}
			b.queue = append(b.queue, &leftNode)
			node.Left = &leftNode
		}

		if val := b.nextValue(); val != nil {
			rightNode := TreeNode{Val: *val}
			b.queue = append(b.queue, &rightNode)
			node.Right = &rightNode
		}
	}

	return &root
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	return &TreeNode{
		Val:   t1.Val + t2.Val,
		Left:  mergeTrees(t1.Left, t2.Left),
		Right: mergeTrees(t1.Right, t2.Right),
	}
}
