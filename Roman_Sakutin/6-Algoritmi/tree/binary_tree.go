package main
//
//type Node struct {
//	Value int
//	Left  *Node
//	Right *Node
//}
//
//func NewNode(value int) *Node {
//	return &Node{
//		Value: value,
//		Left:  nil,
//		Right: nil,
//	}
//}
//
//func (n *Node) insert(value int) *Node {
//	if n == nil {
//		return NewNode(value)
//	}
//	if value > n.Value {
//		n.Right = n.Right.insert(value)
//	} else if value < n.Value {
//		n.Left = n.Left.insert(value)
//	}
//	return n
//}
//
//func (n *Node) search(value int) *Node {
//	if n == nil {
//		return nil
//	}
//	if value == n.Value {
//		return n
//	}
//	if value > n.Value {
//		return n.Right.search(value)
//	} else {
//		return n.Left.search(value)
//	}
//}
//
//func (n *Node) remove(value int) *Node {
//	if n == nil {
//		return nil
//	}
//	if value == n.Value {
//		if n.Left != nil && n.Right != nil {
//			mini := n.Right.minValue()
//			n.Value = mini.Value
//			n.Right = n.Right.remove(mini.Value)
//			return n
//		}
//		if n.Right != nil {
//			return n.Right
//		} else if n.Left != nil {
//			return n.Left
//		}
//		return nil
//	}
//	if value > n.Value {
//		n.Right = n.Right.remove(value)
//	} else {
//		n.Left = n.Left.remove(value)
//	}
//	return n
//}
//
//func (n *Node) minValue() *Node {
//	if n.Left == nil {
//		return n
//	}
//	return n.Left.minValue()
//}
//
//func (n *Node) maxValue() *Node {
//	if n.Right == nil {
//		return n
//	}
//	return n.Right.maxValue()
//}
