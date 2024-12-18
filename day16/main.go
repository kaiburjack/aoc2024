package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
	"time"
)

type coord struct {
	x, y, d int
}

type node struct {
	neighbors   [3]*node
	prev        []*node
	c           coord
	cost, dcost uint
	end, seen   bool
}

const rotateCost = 1000
const uintmax = ^uint(0)

func isWalkable(x, y int, grid [][]byte) bool {
	return x >= 0 && y >= 0 && y < len(grid) && x < len(grid[y]) && grid[y][x] != '#'
}

var d2dd = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
var dirCost = [3]uint{rotateCost + 1, 1, rotateCost + 1}

// buildGraph builds a graph where each node is a path within the grid
// without any "decision points" - that is, that path can be walked (possibly through corners)
// without encountering a junction.
// The node returned by any particular call to buildGraph is the node representing
// the path that starts at the given coordinates and direction.
// If the node would lead to a dead-end, this function returns nil.
func buildGraph(cx, cy, d int, cost uint, grid [][]byte, seen map[coord]*node) *node {
	if v, ok := seen[coord{cx, cy, d}]; ok {
		return v
	}
	n := &node{c: coord{cx, cy, d}}
	seen[coord{cx, cy, d}] = n
	for {
		var walkable [3]bool
		if grid[cy][cx] == 'E' {
			n.cost, n.dcost, n.end = cost, uintmax, true
			return n
		}
		wc, wi := 0, 0
		for i := -1; i <= 1; i++ {
			nd := d2dd[(d+4+i)&3]
			nx, ny := cx+nd[0], cy+nd[1]
			if isWalkable(nx, ny, grid) {
				walkable[i+1], wc, wi = true, wc+1, i
			}
		}
		if wc == 0 {
			return nil
		} else if wc == 1 {
			nd := (d + 4 + wi) & 3
			ndd := d2dd[nd]
			if wi == 0 {
				cx, cy, d, cost = cx+ndd[0], cy+ndd[1], nd, cost+dirCost[wi+1]
			} else {
				n.cost, n.dcost = cost, uintmax
				if ne := buildGraph(cx+ndd[0], cy+ndd[1], nd, 0, grid, seen); ne != nil {
					n.neighbors[wi+1] = ne
				} else {
					return nil
				}
				return n
			}
		} else {
			n.cost, n.dcost = cost, uintmax
			for i := 0; i < 3; i++ {
				if !walkable[i] {
					continue
				}
				nd := (d + 4 + i - 1) & 3
				ndd := d2dd[nd]
				if ne := buildGraph(cx+ndd[0], cy+ndd[1], nd, 0, grid, seen); ne != nil {
					n.neighbors[i] = ne
				}
			}
			return n
		}
	}
}

type priorityQueue []*node

func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].dcost < pq[j].dcost
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*node))
}
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func dijkstra(start *node) (uint, uint64) {
	start.dcost = start.cost
	var pq priorityQueue
	heap.Push(&pq, start)
	minCostToEnd := uintmax
	var ends []*node
	for pq.Len() > 0 {
		n := heap.Pop(&pq).(*node)
		if n.end {
			if n.dcost < minCostToEnd {
				ends = append(ends, n)
				minCostToEnd = n.dcost
			}
			continue
		}
		n.seen = true
		for i, ne := range n.neighbors {
			if ne == nil || ne.seen {
				continue
			}
			c := n.dcost + ne.cost + dirCost[i]
			if c < ne.dcost {
				ne.dcost = c
				ne.prev = []*node{n}
				heap.Push(&pq, ne)
			} else if c == ne.dcost {
				ne.prev = append(ne.prev, n)
			}
		}
	}
	return minCostToEnd, computeNumberOfDistinctPointsVisited(ends[0])
}

func computeNumberOfDistinctPointsVisited(end *node) uint64 {
	var count uint64
	var queue []*node
	queue = append(queue, end)
	seen := make(map[coord]struct{})
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		// compute all positions that can be reached from this node
		nd := d2dd[n.c.d]
		for i := 0; i < int(n.cost)+1; i++ {
			nx, ny := n.c.x+nd[0]*i, n.c.y+nd[1]*i
			if _, ok := seen[coord{nx, ny, 0}]; ok {
				continue
			}
			seen[coord{nx, ny, 0}] = struct{}{}
			count++
		}
		for _, ne := range n.prev {
			queue = append(queue, ne)
		}
	}
	return count
}

func input() ([][]byte, int, int) {
	file, _ := os.Open("real.txt")
	scanner := bufio.NewScanner(file)
	var grid [][]byte
	sx, sy := 0, 0
	for scanner.Scan() {
		row := scanner.Text()
		if idx := strings.IndexByte(row, 'S'); idx != -1 {
			sx, sy = idx, len(grid)
		}
		grid = append(grid, []byte(row))
	}
	return grid, sx, sy
}

func main() {
	grid, sx, sy := input()
	seen := make(map[coord]*node)
	time1 := time.Now()
	start := buildGraph(sx, sy, 0, 0, grid, seen)
	result, uniquePoints := dijkstra(start)
	time2 := time.Now()
	fmt.Printf("Part1: %v\n", result)
	fmt.Printf("Part2: %v\n", uniquePoints)
	fmt.Printf("Total time: %v\n", time2.Sub(time1))
}
