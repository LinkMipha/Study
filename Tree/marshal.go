package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

const SEP = ","
const NULL = "#"

type Codec struct {}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var queue []*TreeNode
	queue = append(queue, root)
	var list []string
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			list = append(list, NULL)
			continue
		}
		list = append(list, strconv.Itoa(node.Val))
		queue = append(queue, node.Left, node.Right)
	}
	return strings.Join(list, SEP)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, SEP)
	if nodes[0] == "#" {
		return nil
	}
	root := new(TreeNode)
	root.Val, _ = strconv.Atoi(nodes[0])
	var queue []*TreeNode
	queue = append(queue, root)
	for i := 0; len(queue) > 0; {
		parent := queue[0]
		queue = queue[1:]
		i++
		left := nodes[i]
		if left != NULL {
			node := new(TreeNode)
			node.Val, _ = strconv.Atoi(left)
			parent.Left = node
			queue = append(queue, parent.Left)
		}
		i++
		right := nodes[i]
		if right != NULL {
			node := new(TreeNode)
			node.Val, _ = strconv.Atoi(right)
			parent.Right = node
			queue = append(queue, parent.Right)
		}
	}
	return root
}

