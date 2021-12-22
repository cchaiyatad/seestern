package random

import (
	"fmt"
)

type (
	Any       interface{}
	LessFunc  func(interface{}, interface{}) bool
	GetFunc   func(interface{}, interface{}) bool
	VisitFunc func(interface{}) bool

	node struct {
		value               Any
		parent, left, right *node
	}
	nodei struct {
		step int
		node *node
		prev *nodei
	}

	SplayTree struct {
		length int
		root   *node
		less   LessFunc
		get    GetFunc
	}
)

// Create a new splay tree, using the less function to determine the order.
func NewSplayTree(less LessFunc, get GetFunc) *SplayTree {
	return &SplayTree{0, nil, less, get}
}

// Get the first value from the collection. Returns nil if empty.
func (tree *SplayTree) First() Any {
	if tree.length == 0 {
		return nil
	}

	n := tree.root
	for n.left != nil {
		n = n.left
	}
	return n.value
}

// Get the last value from the collection. Returns nil if empty.
func (tree *SplayTree) Last() Any {
	if tree.length == 0 {
		return nil
	}
	n := tree.root
	for n.right != nil {
		n = n.right
	}
	return n.value
}

// Get an item from the splay tree
func (tree *SplayTree) Get(item Any) Any {
	if tree.length == 0 {
		return nil
	}

	n := tree.root
	for n != nil {
		if tree.get(item, n.value) {
			tree.splay(n)
			return n.value
		}

		if tree.less(item, n.value) {
			n = n.left
			continue
		}

		if tree.less(n.value, item) {
			n = n.right
			continue
		}

	}
	return nil
}
func (tree *SplayTree) Has(value Any) bool {
	return tree.Get(value) != nil
}
func (tree *SplayTree) Init() {
	tree.length = 0
	tree.root = nil
}
func (tree *SplayTree) Add(value Any) {
	if tree.length == 0 {
		tree.root = &node{value, nil, nil, nil}
		tree.length = 1
		return
	}

	n := tree.root
	for {
		if tree.less(value, n.value) {
			if n.left == nil {
				n.left = &node{value, n, nil, nil}
				tree.length++
				n = n.left
				break
			}
			n = n.left
			continue
		}

		if tree.less(n.value, value) {
			if n.right == nil {
				n.right = &node{value, n, nil, nil}
				tree.length++
				n = n.right
				break
			}
			n = n.right
			continue
		}

		n.value = value
		break
	}
	tree.splay(n)
}
func (tree *SplayTree) PreOrder(visit VisitFunc) {
	if tree.length == 1 {
		return
	}
	i := &nodei{0, tree.root, nil}

traversal:
	for i != nil {
		switch i.step {
		// Value
		case 0:
			i.step++
			if !visit(i.node.value) {
				break traversal
			}
		// Left
		case 1:
			i.step++
			if i.node.left != nil {
				i = &nodei{0, i.node.left, i}
			}
		// Right
		case 2:
			i.step++
			if i.node.right != nil {
				i = &nodei{0, i.node.right, i}
			}
		default:
			i = i.prev
		}
	}
}
func (tree *SplayTree) InOrder(visit VisitFunc) {
	if tree.length == 1 {
		return
	}
	i := &nodei{0, tree.root, nil}

traversal:
	for i != nil {
		switch i.step {
		// Left
		case 0:
			i.step++
			if i.node.left != nil {
				i = &nodei{0, i.node.left, i}
			}
		// Value
		case 1:
			i.step++
			if !visit(i.node.value) {
				break traversal
			}
		// Right
		case 2:
			i.step++
			if i.node.right != nil {
				i = &nodei{0, i.node.right, i}
			}
		default:
			i = i.prev
		}
	}
}
func (tree *SplayTree) PostOrder(visit VisitFunc) {
	if tree.length == 1 {
		return
	}
	i := &nodei{0, tree.root, nil}
traversal:
	for i != nil {
		switch i.step {
		// Left
		case 0:
			i.step++
			if i.node.left != nil {
				i = &nodei{0, i.node.left, i}
			}
		// Right
		case 1:
			i.step++
			if i.node.right != nil {
				i = &nodei{0, i.node.right, i}
			}
		// Value
		case 2:
			i.step++
			if !visit(i.node.value) {
				break traversal
			}
		default:
			i = i.prev
		}
	}
}
func (tree *SplayTree) Do(visit VisitFunc) {
	tree.InOrder(visit)
}
func (tree *SplayTree) Len() int {
	return tree.length
}
func (tree *SplayTree) Remove(value Any) {
	if tree.length == 0 {
		return
	}

	n := tree.root
	for n != nil {
		if tree.less(value, n.value) {
			n = n.left
			continue
		}
		if tree.less(n.value, value) {
			n = n.right
			continue
		}

		// First splay the parent node
		if n.parent != nil {
			tree.splay(n.parent)
		}

		// No children
		if n.left == nil && n.right == nil {
			// guess we're the root node
			if n.parent == nil {
				tree.root = nil
				break
			}
			if n.parent.left == n {
				n.parent.left = nil
			} else {
				n.parent.right = nil
			}
		} else if n.left == nil {
			// root node
			if n.parent == nil {
				tree.root = n.right
				break
			}
			if n.parent.left == n {
				n.parent.left = n.right
			} else {
				n.parent.right = n.right
			}
		} else if n.right == nil {
			// root node
			if n.parent == nil {
				tree.root = n.left
				break
			}
			if n.parent.left == n {
				n.parent.left = n.left
			} else {
				n.parent.right = n.left
			}
		} else {
			// find the successor
			s := n.right
			for s.left != nil {
				s = s.left
			}

			np := n.parent
			nl := n.left
			nr := n.right

			sp := s.parent
			sr := s.right

			// Update parent
			s.parent = np
			if np == nil {
				tree.root = s
			} else {
				if np.left == n {
					np.left = s
				} else {
					np.right = s
				}
			}

			// Update left
			s.left = nl
			s.left.parent = s

			// Update right
			if nr != s {
				s.right = nr
				s.right.parent = s
			}

			// Update successor parent
			if sp.left == s {
				sp.left = sr
			} else {
				sp.right = sr
			}
		}

		break
	}

	if n != nil {
		tree.length--
	}
}
func (tree *SplayTree) String() string {
	if tree.length == 0 {
		return "{}"
	}
	return tree.root.String()
}

// Splay a node in the tree (send it to the top)
func (tree *SplayTree) splay(n *node) {
	// Already root, nothing to do
	if n.parent == nil {
		tree.root = n
		return
	}

	p := n.parent
	g := p.parent

	// Zig
	if p == tree.root {
		if n == p.left {
			p.rotateRight()
		} else {
			p.rotateLeft()
		}
	} else {
		// Zig-zig
		if n == p.left && p == g.left {
			g.rotateRight()
			p.rotateRight()
		} else if n == p.right && p == g.right {
			g.rotateLeft()
			p.rotateLeft()
			// Zig-zag
		} else if n == p.right && p == g.left {
			p.rotateLeft()
			g.rotateRight()
		} else if n == p.left && p == g.right {
			p.rotateRight()
			g.rotateLeft()
		}
	}
	tree.splay(n)
}

// Node methods
func (tree *node) String() string {
	str := "{" + fmt.Sprint(tree.value) + "|"
	if tree.left != nil {
		str += tree.left.String()
	}
	str += "|"
	if tree.right != nil {
		str += tree.right.String()
	}
	str += "}"
	return str
}
func (tree *node) rotateLeft() {
	parent := tree.parent
	pivot := tree.right
	if pivot == nil {
		return
	}
	child := pivot.left

	// Update the parent
	if parent != nil {
		if parent.left == tree {
			parent.left = pivot
		} else {
			parent.right = pivot
		}
	}

	// Update the pivot
	pivot.parent = parent
	pivot.left = tree

	// Update the child
	if child != nil {
		child.parent = tree
	}

	// Update tree
	tree.parent = pivot
	tree.right = child
}
func (tree *node) rotateRight() {
	parent := tree.parent
	pivot := tree.left
	if pivot == nil {
		return
	}
	child := pivot.right

	// Update the parent
	if parent != nil {
		if parent.left == tree {
			parent.left = pivot
		} else {
			parent.right = pivot
		}
	}

	// Update the pivot
	pivot.parent = parent
	pivot.right = tree

	if child != nil {
		child.parent = tree
	}

	// Update tree
	tree.parent = pivot
	tree.left = child
}
