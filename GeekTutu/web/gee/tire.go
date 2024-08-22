package gee

type node struct {
	pattern  string // 整个curl
	part     string // 当前部分
	children []*node
	vague    bool // 是否模糊匹配
}

func (n *node) findChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.vague {
			return child
		}
	}

	return nil
}

// 匹配某个节点的所有子节点，把满足条件的放在slice里返回
func (n *node) matchChildren(part string) []*node {
	var x []*node
	for _, child := range n.children {
		if child.part == part || child.vague {
			x = append(x, child)
		}
	}

	return x
}

// /p/s/a parts = [p,s,a] 用part 去匹配，没有就新建一个node，并操作这个node去递归
//
//	part = parts[height] 递归时 height + 1
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.findChild(part)
	if child == nil {
		child = &node{part: part, vague: part[0] == ':' || part[0] == '*'}
		n.children = append(child.children, child)
	}

	child.insert(pattern, parts, height+1)
}

// 匹配路由，找不到就是nil
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height {
		if n.pattern == "" {
			return nil
		}

		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
