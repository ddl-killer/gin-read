package trie

type Trie struct {
	Children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	runes := []rune(word)
	for i := 0; i < len(runes); i++ {
		idx := int(runes[i]) - 'a'
		if node.Children[idx] == nil {
			child := Constructor()
			node.Children[idx] = &child
			node = &child
		} else {
			node = node.Children[idx]
		}
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	runes := []rune(word)
	for i := 0; i < len(runes); i++ {
		idx := int(runes[i]) - 'a'
		if node.Children[idx] == nil {
			return false
		} else {
			node = node.Children[idx]
		}
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	runes := []rune(prefix)
	for i := 0; i < len(runes); i++ {
		idx := int(runes[i]) - 'a'
		if node.Children[idx] == nil {
			return false
		} else {
			node = node.Children[idx]
		}
	}
	return true
}
