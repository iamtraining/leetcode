package main

func main() {

}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	return helper(node, map[int]*Node{})
}

func helper(node *Node, visited map[int]*Node) *Node {
	if visited[node.Val] != nil {
		return visited[node.Val]
	}

	new := &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}

	visited[node.Val] = new
	if node.Neighbors != nil {
		for _, n := range node.Neighbors {
			new.Neighbors = append(new.Neighbors, helper(n, visited))
		}
	}
	return new
}
