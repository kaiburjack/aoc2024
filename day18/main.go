package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type coord struct {
	x, y int
}

type node struct {
	neighbors [4]*node
	dcost     uint
	end, seen bool
}

func isWalkable(x, y int, grid [ey + 1][ex + 1]byte) bool {
	return x >= 0 && y >= 0 && y <= ey && x <= ex && grid[y][x] != '#'
}

var d2dd = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

const ex = 70
const ey = 70

const maxuint = ^uint(0)

func buildGraph(cx, cy int, grid [ey + 1][ex + 1]byte, seen map[coord]*node) *node {
	if v, ok := seen[coord{cx, cy}]; ok {
		return v
	}
	n := &node{dcost: maxuint}
	if cx == ex && cy == ey {
		n.end = true
		return n
	}
	seen[coord{cx, cy}] = n
	for i := 0; i < 4; i++ {
		nx, ny := cx+d2dd[i][0], cy+d2dd[i][1]
		if !isWalkable(nx, ny, grid) {
			continue
		}
		if v, ok := seen[coord{nx, ny}]; ok {
			n.neighbors[i] = v
			continue
		}
		n.neighbors[i] = buildGraph(nx, ny, grid, seen)
	}
	return n
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

func input(n int) ([ey + 1][ex + 1]byte, int, int) {
	file, _ := os.Open("real.txt")
	scanner := bufio.NewScanner(file)
	var grid [ey + 1][ex + 1]byte
	// fill grid with dots
	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = '.'
		}
	}
	bx, by := 0, 0
	for i := 0; i < n && scanner.Scan(); i++ {
		row := scanner.Text()
		splitted := strings.Split(row, ",")
		leftAsInt, _ := strconv.Atoi(splitted[0])
		rightAsInt, _ := strconv.Atoi(splitted[1])
		grid[rightAsInt][leftAsInt] = '#'
		bx, by = leftAsInt, rightAsInt
	}
	return grid, bx, by
}

func printGrid(grid [ey + 1][ex + 1]byte) {
	for y := range grid {
		for x := range grid[y] {
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Println()
	}
}

func dijkstra(start *node) uint {
	start.dcost = 0
	pq := &priorityQueue{start}
	start.seen = true
	for pq.Len() > 0 {
		n := heap.Pop(pq).(*node)
		if n.end {
			return n.dcost
		}
		for _, neighbor := range n.neighbors {
			if neighbor == nil {
				continue
			}
			if neighbor.seen {
				continue
			}
			neighbor.seen = true
			c := n.dcost + 1
			if c < neighbor.dcost {
				neighbor.dcost = n.dcost + 1
				heap.Push(pq, neighbor)
			}
		}
	}
	return maxuint
}

func main() {
	// part 1
	{
		grid, _, _ := input(1024)
		seen := make(map[coord]*node)
		time1 := time.Now()
		start := buildGraph(0, 0, grid, seen)
		dcost := dijkstra(start)
		time2 := time.Now()
		fmt.Printf("Start: %v\n", start)
		fmt.Printf("Distance: %d\n", dcost)
		fmt.Printf("Time: %v\n", time2.Sub(time1))
	}
	// part 2
	{
		time1 := time.Now()
		for i := 1025; i < 1000000; i++ {
			grid, bx, by := input(i)
			seen := make(map[coord]*node)
			start := buildGraph(0, 0, grid, seen)
			dcost := dijkstra(start)
			if dcost == maxuint {
				fmt.Printf("Part 2: %d, %d\n", bx, by)
				break
			}
		}
		time2 := time.Now()
		fmt.Printf("Time: %v\n", time2.Sub(time1))
	}
}
