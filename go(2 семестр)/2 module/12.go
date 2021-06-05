package main

import (
	"container/heap"
	"fmt"

	"github.com/skorobogatov/input"
)

var G []*Vertex

type Vertex struct {
	order  int
	val    int
	i      int
	j      int
	dist   int
	parent *Vertex
}

func Relax(u *Vertex, i int) bool {
	v := G[i]
	changed := (u.dist+v.val < v.dist)
	if changed {
		v.dist = u.dist + v.val
		v.parent = u
	}
	return changed
}

type PriorityQueue []*Vertex

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(a, b int) bool {
	return pq[a].dist < pq[b].dist
}

func (pq PriorityQueue) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
	pq[a].order = a
	pq[b].order = b
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	vertex := x.(*Vertex)
	vertex.order = n
	*pq = append(*pq, vertex)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	vertex := old[n-1]
	vertex.order = -1
	*pq = old[0 : n-1]
	return vertex
}

func InitPriorityQueue(root *Vertex) PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, root)
	return pq
}

func Dijcstra(n int) {
	pq := InitPriorityQueue(G[0])
	for {
		if pq.Len() == 0 {
			break
		}
		v := heap.Pop(&pq).(*Vertex)
		if v.i > 0 {
			temp := (v.i-1)*n + v.j
			if Relax(v, temp) {
				heap.Push(&pq, G[temp])
			}
		}
		if v.i < n-1 {
			temp := (v.i+1)*n + v.j
			if Relax(v, temp) {
				heap.Push(&pq, G[temp])
			}
		}

		if v.j > 0 {
			temp := v.i*n + v.j - 1
			if Relax(v, temp) {
				heap.Push(&pq, G[temp])
			}
		}
		if v.j < n-1 {
			temp := v.i*n + v.j + 1
			if Relax(v, temp) {
				heap.Push(&pq, G[temp])
			}
		}

	}
}
func create(val, i, j, dist, n int) Vertex {
	var cur Vertex
	cur.dist = dist
	cur.i = i
	cur.j = j
	cur.val = val
	cur.order = i*n + j
	return cur
}
func main() {
	var n int
	input.Scanf("%d", &n)
	G = make([]*Vertex, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var val int
			input.Scanf("%d", &val)
			if i+j == 0 {
				temp := create(val, i, j, val, n)
				G[i*n+j] = &temp
			} else {
				temp := create(val, i, j, 10000, n)
				G[i*n+j] = &temp
			}

		}
	}
	Dijcstra(n)
	last := G[n*n-1]
	fmt.Println(last.dist)
}

