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
	neighbors [3]*node
	cost      uint64
	dcost     uint64
	deadEnd   bool
	end       bool
}

const rotateCost = 1000
const uint64max = ^uint64(0)

func isWalkable(x, y int, grid [][]byte) bool {
	return x >= 0 && y >= 0 && y < len(grid) && x < len(grid[y]) && grid[y][x] != '#'
}

var d2dd = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
var dirCost = [3]uint64{rotateCost + 1, 1, rotateCost + 1}

// buildGraph builds a graph where each node is a path within the grid
// without any "decision points" - that is, that path can be walked (possibly through corners)
// without encountering a junction.
// The node returned by any particular call to buildGraph is the node representing
// the path that starts at the given coordinates and direction.
// If the node would lead to a dead-end, this function returns nil.
func buildGraph(cx, cy, d int, cost uint64, grid [][]byte, seen map[coord]*node) *node {
	if v, ok := seen[coord{cx, cy, d}]; ok {
		return v
	}
	n := &node{deadEnd: true}
	seen[coord{cx, cy, d}] = n
	for {
		var walkable [3]bool
		if grid[cy][cx] == 'E' {
			n.cost, n.dcost, n.end, n.deadEnd = cost, uint64max, true, false
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
			cx, cy, d, cost = cx+ndd[0], cy+ndd[1], nd, cost+dirCost[wi+1]
		} else {
			n.cost, n.dcost, n.deadEnd = cost, uint64max, false
			for i := 0; i < 3; i++ {
				if !walkable[i] {
					continue
				}
				nd := (d + 4 + i - 1) & 3
				ndd := d2dd[nd]
				if ne := buildGraph(cx+ndd[0], cy+ndd[1], nd, 0, grid, seen); ne != nil && !ne.deadEnd {
					n.neighbors[i] = ne
				}
			}
			allDeadEnd := true
			for i := 0; i < 3 && allDeadEnd; i++ {
				if n.neighbors[i] != nil && !n.neighbors[i].deadEnd {
					allDeadEnd = false
				}
			}
			if allDeadEnd {
				n.deadEnd = true
				return nil
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

func dijkstra(start *node) uint64 {
	start.dcost = start.cost
	pq, seen := make(priorityQueue, 0, 8), make(map[*node]struct{})
	heap.Push(&pq, start)
	for pq.Len() > 0 {
		n := heap.Pop(&pq).(*node)
		if n.end {
			return n.dcost
		}
		if _, ok := seen[n]; ok {
			continue
		}
		seen[n] = struct{}{}
		for i, ne := range n.neighbors {
			if ne == nil || ne.deadEnd {
				continue
			}
			if _, ok := seen[ne]; ok {
				continue
			}
			if n.dcost+ne.cost+dirCost[i] < ne.dcost {
				ne.dcost = n.dcost + ne.cost + dirCost[i]
				heap.Push(&pq, ne)
			}
		}
	}
	panic("should not happen")
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
	// part 1
	{
		time1 := time.Now()
		start := buildGraph(sx, sy, 0, 0, grid, seen)
		result := dijkstra(start)
		time2 := time.Now()
		fmt.Printf("Part1 (in %v): %v\n", time2.Sub(time1), result)
	}
}
