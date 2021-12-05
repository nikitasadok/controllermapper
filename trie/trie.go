package trie

type Node struct {
	Children     map[rune]*Node
	Value        rune
	IsController bool
}

func (n *Node) Find(node *Node, key string) bool {
	var lastController string
	convKey := []rune(key)
	for i := 0; i < len(convKey); i++ {
		next, ok := node.Children[convKey[i]]
		if !ok {
			return lastController != ""
		}
		if next.IsController {
			lastController = string(convKey[:i])
		}
		node = next
	}

	return lastController != ""
}

func (n *Node) Insert(key string) {
	cur := n
	convKey := []rune(key)
	for i := 0; i < len(convKey); i++ {
		_, ok := cur.Children[convKey[i]]
		if !ok {
			cur.Children[convKey[i]] = &Node{Children: map[rune]*Node{}}
		}
		cur = cur.Children[convKey[i]]
	}
	cur.IsController = true
}
