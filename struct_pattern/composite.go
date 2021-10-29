package struct_pattern

import "fmt"

type ChainNode interface {
	GetName() string
	SetName(name string)
	GetMainNode() ChainNode
	SetMainNode(node ChainNode)
	AddChild(node ChainNode)
	Print(segment string)
}

type Node struct {
	parent ChainNode
	name string
}

func (n *Node) GetName() string {
	return n.name
}

func (n *Node) SetName(name string) {
		n.name = name
}

func (n *Node) GetMainNode() ChainNode {
	return n.parent
}

func (n *Node) SetMainNode(node ChainNode) {
	n.parent = node
}

func (n *Node) AddChild(node ChainNode) {}

func (n *Node) Print(segment string) {
	fmt.Printf("%s-%s\n", segment, n.name)
}

type Leaf struct {
	Node
}

func NewLeaf(name string) ChainNode {
	return &Leaf{
		Node{
			name: name,
		},
	}
}

type NodeStruct struct {
	Node
	leaf []ChainNode
}

func NewNodeStruct(name string) ChainNode {
	return &NodeStruct{
		Node: Node{
			name: name,
		},
	}
}

func (n *NodeStruct) GetName() string {
	return n.name
}

func (n *NodeStruct) SetName(name string) {
	n.name = name
}

func (n *NodeStruct) GetMainNode() ChainNode {
	return n.parent
}

func (n *NodeStruct) SetMainNode(node ChainNode) {
	n.parent = node
}

func (n *NodeStruct) AddChild(node ChainNode) {
	n.leaf = append(n.leaf, node)
}

func (n *NodeStruct) Print(segment string) {
	fmt.Printf("%s+%s\n", segment, n.GetName())
	segment += " "
	for _, node := range n.leaf {
		node.Print(segment)
	}
}
