package struct_pattern

import (
	"testing"
)

func Test_Composite(t *testing.T) {
	root := NewNodeStruct("root")
	c1 := NewNodeStruct("c1")
	c2 := NewNodeStruct("c2")
	c3 := NewNodeStruct("c3")

	l1 := NewLeaf("l1")
	l2 := NewLeaf("l2")
	l3 := NewLeaf("l3")

	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(c3)
	c1.AddChild(l1)
	c1.AddChild(l2)
	c1.AddChild(l3)

	root.Print("")
}
