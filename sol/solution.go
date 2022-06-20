package sol

func findRedundantConnection(edges [][]int) []int {
	parent := make(map[int]int)
	rank := make(map[int]int)
	nLen := len(edges)
	for node := 1; node <= nLen; node++ {
		parent[node] = node
		rank[node] = 1
	}
	var find func(node int) int
	find = func(node int) int {
		p := parent[node]
		for p != parent[p] {
			parent[p] = parent[parent[p]]
			p = parent[p]
		}
		return p
	}
	var union func(node1, node2 int) bool
	union = func(node1, node2 int) bool {
		p1 := find(node1)
		p2 := find(node2)
		if p1 == p2 {
			return false
		}
		if rank[p1] > rank[p2] {
			parent[p2] = p1
			rank[p1] += rank[p2]
		} else {
			parent[p1] = p2
			rank[p2] += rank[p1]
		}
		return true
	}
	for _, nodes := range edges {
		if !union(nodes[0], nodes[1]) {
			return nodes
		}
	}
	return []int{}
}
