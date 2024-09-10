package tree

import "fmt"

type SearchBinaryTree struct {
	Root      *Node
	Ligations map[string]string
}

func NewSearchBinaryTree() *SearchBinaryTree {
	return &SearchBinaryTree{
		Root:      nil,
		Ligations: make(map[string]string),
	}
}

func (s *SearchBinaryTree) Insert(value int) {
	newNode := NewNode(value)
	if s.Root == nil {
		s.Root = newNode
	} else {
		actual := s.Root
		for {
			father := actual
			if value == actual.Value {
				return
			} else if value < actual.Value {
				actual = actual.Left
				if actual == nil {
					father.Left = newNode
					s.Ligations[fmt.Sprintf("%d->%d", father.Value, newNode.Value)] = fmt.Sprintf("%d->%d", father.Value, newNode.Value)
					return
				}
			} else {
				actual = actual.Right
				if actual == nil {
					father.Right = newNode
					s.Ligations[fmt.Sprintf("%d->%d", father.Value, newNode.Value)] = fmt.Sprintf("%d->%d", father.Value, newNode.Value)
					return
				}
			}
		}
	}
}

func (s *SearchBinaryTree) Search(value int) *Node {
	actual := s.Root
	for {
		if actual.Value == value {
			return actual
		}

		if value < actual.Value {
			actual = actual.Left
		} else {
			actual = actual.Right
		}

		if actual == nil {
			return nil
		}
	}

}

func (s *SearchBinaryTree) PreOrder(node *Node) {
	if node != nil {
		fmt.Println(node.Value)
		s.PreOrder(node.Left)
		s.PreOrder(node.Right)
	}
}

func (s *SearchBinaryTree) InOrder(node *Node) {
	if node != nil {
		s.InOrder(node.Left)
		fmt.Println(node.Value)
		s.InOrder(node.Right)
	}
}

func (s *SearchBinaryTree) PosOrder(node *Node) {
	if node != nil {
		s.InOrder(node.Left)
		s.InOrder(node.Right)
		fmt.Println(node.Value)

	}
}

func (s *SearchBinaryTree) Delete(value int) {
	if s.Root == nil {
		fmt.Println("tree is empty")
		return
	}
	actual := s.Root
	isLeft := true
	for {
		if actual.Value == value {
			return
		}
		father := actual
		if value < actual.Value {
			isLeft = true
			actual = actual.Left
		} else {
			isLeft = false
			actual = actual.Right
		}
		if actual == nil {
			return
		}
		if actual.Left == nil && actual.Right == nil {
			if actual == s.Root {
				s.Root = nil
			} else if isLeft {
				delete(s.Ligations, fmt.Sprintf("%d->%d", father.Value, actual.Value))
				father.Left = nil
			} else {
				delete(s.Ligations, fmt.Sprintf("%d->%d", father.Value, actual.Value))
				father.Right = nil
			}
		} else if actual.Right == nil {
			delete(s.Ligations, fmt.Sprintf("%d->%d", father.Value, actual.Value))
			delete(s.Ligations, fmt.Sprintf("%d->%d", actual.Value, actual.Left.Value))
			if actual == s.Root {
				s.Root = actual.Left
				s.Ligations[fmt.Sprintf("%d->%d", s.Root.Value, actual.Left.Value)] = fmt.Sprintf("%d->%d", s.Root.Value, actual.Left.Value)

			} else if isLeft {
				father.Left = actual.Left
				s.Ligations[fmt.Sprintf("%d->%d", father.Value, actual.Left.Value)] = fmt.Sprintf("%d->%d", father.Value, actual.Left.Value)

			} else {
				father.Right = actual.Left
				s.Ligations[fmt.Sprintf("%d->%d", father.Value, actual.Left.Value)] = fmt.Sprintf("%d->%d", father.Value, actual.Left.Value)
			}
		} else if actual.Left == nil {
			delete(s.Ligations, fmt.Sprintf("%d->%d", father.Value, actual.Value))
			delete(s.Ligations, fmt.Sprintf("%d->%d", actual.Value, actual.Right.Value))
			if actual == s.Root {
				s.Ligations[fmt.Sprintf("%d->%d", s.Root.Value, actual.Right.Value)] = fmt.Sprintf("%d->%d", s.Root.Value, actual.Right.Value)
				s.Root = actual.Right

			} else if isLeft {

				s.Ligations[fmt.Sprintf("%d->%d", father.Value, actual.Right.Value)] = fmt.Sprintf("%d->%d", father.Value, actual.Right.Value)
				father.Left = actual.Right

			} else {
				s.Ligations[fmt.Sprintf("%d->%d", father.Value, actual.Right.Value)] = fmt.Sprintf("%d->%d", father.Value, actual.Right.Value)
				father.Right = actual.Right
			}
		}
	}

}
