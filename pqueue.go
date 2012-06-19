package astar

import (
	"container/heap"
	//"fmt"
)


// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Node


/* sort.Interface */
func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use smaller than here.
    return pq[i].f < pq[j].f
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].heap_index = i
    pq[j].heap_index = j
}

/* heap.interface */
func (pq *PriorityQueue) Push(x interface{}) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    // To simplify indexing expressions in these methods, we save a copy of the
    // slice object. We could instead write (*pq)[i].
    a := *pq
    n := len(a)
    a = a[0 : n+1]
    item := x.(*Node)
    item.heap_index = n
    a[n] = item
    *pq = a
}

func (pq *PriorityQueue) Pop() interface{} {
    a := *pq
    n := len(a)
    item := a[n-1]
    item.heap_index = -1 // for safety
    *pq = a[0 : n-1]
    return item
}



/* Node */

func (pq *PriorityQueue) PushNode(n *Node) {
		heap.Push(pq, n)
}

func (pq *PriorityQueue) PopNode() (*Node) {
		return heap.Pop(pq).(*Node)
}


func (pq *PriorityQueue) RemoveNode(n *Node) {
	heap.Remove(pq, n.heap_index)
}
