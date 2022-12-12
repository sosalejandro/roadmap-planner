package value_objects

import "fmt"

type BaseNode interface {
	String() string
}

// Node is a representation of an entity and all its possible relationships
type Node struct {
	name string
	//Has     []*Node
	//Creates []*Node
	//Blocks  []*Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%s", n.name)
}

func NewNode(name string) *Node {
	return &Node{name: name}
}
