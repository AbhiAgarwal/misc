// http://rosettacode.org/wiki/Dijkstra%27s_algorithm#Go
package graph

// edge struct holds the bare data needed to define a graph.
type Edge struct {
	vert1, vert2 string
	dist         int
}

// node and neighbor structs hold data useful for the heap-optimized
// Dijkstra's shortest path algorithm
type Node struct {
	vert string     // vertex name
	tent int        // tentative distance
	prev *Node      // previous node in shortest path back to start
	done bool       // true when tent and prev represent shortest path
	nbs  []Neighbor // edges from this vertex
	rx   int        // heap.Remove index
}

type Neighbor struct {
	nd   *Node // node corresponding to vertex
	dist int   // distance to this node (from whatever node references this)
}

// return type
type Path struct {
	path   []string
	length int
}

// ndList implements container/heap
type ndList []*Node

func (n ndList) Len() int           { return len(n) }
func (n ndList) Less(i, j int) bool { return n[i].tent < n[j].tent }
func (n ndList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
	n[i].rx = i
	n[j].rx = j
}
func (n *ndList) Push(x interface{}) {
	nd := x.(*Node)
	nd.rx = len(*n)
	*n = append(*n, nd)
}
func (n *ndList) Pop() interface{} {
	s := *n
	last := len(s) - 1
	r := s[last]
	*n = s[:last]
	r.rx = -1
	return r
}

// linkGraph constructs a linked representation of the input graph,
// with additional fields needed by the shortest path algorithm.
//
// Return value allNodes will contain all nodes found in the input graph,
// even ones not reachable from the start node.
// Return values startNode, endNode will be nil if the specified start or
// end node names are not found in the graph.
func LinkGraph(graph []Edge, directed bool,
	start, end string) (allNodes []*Node, startNode, endNode *Node) {

	all := make(map[string]*Node)
	// one pass over graph to collect nodes and link neighbors
	for _, e := range graph {
		n1 := all[e.vert1]
		n2 := all[e.vert2]
		// add previously unseen nodes
		if n1 == nil {
			n1 = &Node{vert: e.vert1}
			all[e.vert1] = n1
		}
		if n2 == nil {
			n2 = &Node{vert: e.vert2}
			all[e.vert2] = n2
		}
		// link neighbors
		n1.nbs = append(n1.nbs, Neighbor{n2, e.dist})
		if !directed {
			n2.nbs = append(n2.nbs, Neighbor{n1, e.dist})
		}
	}
	allNodes = make([]*Node, len(all))
	var n int
	for _, nd := range all {
		allNodes[n] = nd
		n++
	}
	return allNodes, all[start], all[end]
}
