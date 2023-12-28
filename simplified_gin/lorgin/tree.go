package lorgin

type methodTree struct {
	method string
	root   *node
}

type methodTrees []methodTree

func (trees methodTrees) get(method string) *node {
	for _, tree := range trees {
		if tree.method == method {
			return tree.root
		}
	}
	return nil
}

type node struct {
	fullPath  string
	indices   string
	childeren []*node
	handlers  HandlersChain
}

func (n *node) addRoute(path string, handlers HandlersChain) {
	runes := []rune(path)
	cur := n
	for i := 0; i < len(runes); i++ {
		found := false
		for j, r := range []rune(cur.indices) {
			if r == runes[i] {
				cur = cur.childeren[j]
				found = true
				break
			}
		}
		if !found {
			cur.indices += string(runes[i])
			child := &node{
				fullPath: "",
				indices:  "",
				handlers: nil,
			}
			cur.childeren = append(cur.childeren, child)
			cur = child
		}
		if i == len(runes)-1 {
			cur.fullPath = path
			cur.handlers = handlers
		}
	}
}

func (n *node) getValue(path string) (value nodeValue) {
	targetNode := n.search(path)
	if targetNode != nil {
		value.handlers = targetNode.handlers
		value.fullPath = targetNode.fullPath
	}
	return
}

func (n *node) search(path string) *node {
	runes := []rune(path)
	cur := n
	for i := 0; i < len(runes); i++ {
		found := false
		for j, r := range []rune(cur.indices) {
			if r == runes[i] {
				cur = cur.childeren[j]
				found = true
				break
			}
		}
		if !found {
			return nil
		}
	}
	return cur
}

type nodeValue struct {
	handlers HandlersChain
	fullPath string
}
