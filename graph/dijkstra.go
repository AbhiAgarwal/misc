// http://rosettacode.org/wiki/Dijkstra%27s_algorithm#Go
package graph

import (
	"container/heap"
	"math"
)

// dijkstra is a heap-enhanced version of Dijkstra's shortest path algorithm.
//
// If endNode is specified, only a single path is returned.
// If endNode is nil, paths to all nodes are returned.
//
// Note input allNodes is needed to efficiently accomplish WP steps 1 and 2.
// This initialization could be done in linkGraph, but is done here to more
// closely follow the WP algorithm.
func Dijkstra(allNodes []*Node, startNode, endNode *Node) (pl []Path) {
	// WP steps 1 and 2.
	for _, nd := range allNodes {
		nd.tent = math.MaxInt32
		nd.done = false
		nd.prev = nil
		nd.rx = -1
	}
	current := startNode
	current.tent = 0
	var unvis ndList

	for {
		// WP step 3: update tentative distances to neighbors
		for _, nb := range current.nbs {
			if nd := nb.nd; !nd.done {
				if d := current.tent + nb.dist; d < nd.tent {
					nd.tent = d
					nd.prev = current
					if nd.rx < 0 {
						heap.Push(&unvis, nd)
					} else {
						heap.Fix(&unvis, nd.rx)
					}
				}
			}
		}
		// WP step 4: mark current node visited, record path and distance
		current.done = true
		if endNode == nil || current == endNode {
			// record path and distance for return value
			distance := current.tent
			// recover path by tracing prev links,
			var p []string
			for ; current != nil; current = current.prev {
				p = append(p, current.vert)
			}
			// then reverse list
			for i := (len(p) + 1) / 2; i > 0; i-- {
				p[i-1], p[len(p)-i] = p[len(p)-i], p[i-1]
			}
			pl = append(pl, Path{p, distance}) // pl is return value
			// WP step 5 (case of end node reached)
			if endNode != nil {
				return
			}
		}
		if len(unvis) == 0 {
			break // WP step 5 (case of no more reachable nodes)
		}
		// WP step 6: new current is node with smallest tentative distance
		current = heap.Pop(&unvis).(*Node)
	}
	return
}
