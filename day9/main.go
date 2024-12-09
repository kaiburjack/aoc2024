package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type sectorList struct {
	first *node
	last  *node
	size  int
}

type node struct {
	id   int
	idx  int
	len  int
	next *node
	prev *node
}

func (thiz *sectorList) toString() string {
	stringBuilder := strings.Builder{}
	for n := thiz.first; n != nil; n = n.next {
		if n.id == -1 {
			stringBuilder.WriteString(strings.Repeat(".", n.len))
		} else {
			stringBuilder.WriteString(strings.Repeat(fmt.Sprintf("%v", n.id), n.len))
		}
	}
	return stringBuilder.String()
}

func (thiz *node) append(n *node) {
	if thiz.next != nil {
		thiz.next.prev = n
	}
	n.next = thiz.next
	n.prev = thiz
	n.idx = thiz.idx + thiz.len
	thiz.next = n
}

func (thiz *sectorList) buildFromBytes(bytes []byte, fullFiles bool) {
	currentId := int64(0)
	for i := 0; i < len(bytes); i++ {
		if i&1 == 0 {
			if fullFiles {
				newNode := &node{id: int(currentId), len: int(bytes[i] - '0')}
				thiz.append(newNode)
			} else {
				for j := 0; j < int(bytes[i]-'0'); j++ {
					newNode := &node{id: int(currentId), len: 1}
					thiz.append(newNode)
				}
			}
			currentId++
		} else {
			if fullFiles {
				newNode := &node{id: -1, len: int(bytes[i] - '0')}
				thiz.append(newNode)
			} else {
				for j := 0; j < int(bytes[i]-'0'); j++ {
					newNode := &node{id: -1, len: 1}
					thiz.append(newNode)
				}
			}
		}
	}
}

func (thiz *sectorList) append(n *node) {
	if thiz.first == nil {
		thiz.first = n
		thiz.last = n
	} else {
		thiz.last.append(n)
		thiz.last = n
	}
	thiz.size += n.len
}

func (thiz *sectorList) findFreeSpace(size int) *node {
	for n := thiz.first; n != nil; n = n.next {
		if n.id == -1 && n.len >= size {
			return n
		}
	}
	return nil
}

func (thiz *sectorList) move(dst *node, src *node) {
	if dst.len == src.len {
		dst.id = src.id
		src.id = -1
	} else if dst.len > src.len {
		// create new node with remaining space
		newNode := &node{id: dst.id, len: dst.len - src.len}
		dst.append(newNode)
		dst.id = src.id
		dst.len = src.len
		src.id = -1
	} else {
		panic("should not happen!")
	}
}

func (thiz *sectorList) compact() {
	for n := thiz.last; n != nil; n = n.prev {
		if n.id == -1 {
			continue
		}
		free := thiz.findFreeSpace(n.len)
		if free != nil {
			if free.idx > n.idx {
				continue
			}
			thiz.move(free, n)
		}
	}
}

func (thiz *sectorList) computeChecksum() int64 {
	checksum := int64(0)
	idx := 0
	for n := thiz.first; n != nil; n = n.next {
		if n.id == -1 {
			idx += n.len
			continue
		}
		for i := 0; i < n.len; i++ {
			checksum += int64(n.id) * int64(idx)
			idx++
		}
	}
	return checksum
}

func main() {
	bytes, _ := os.ReadFile("real.txt")
	// part 1
	{
		time1 := time.Now()
		list := &sectorList{}
		list.buildFromBytes(bytes, false)
		list.compact()
		time2 := time.Now()
		fmt.Printf("Checksum: %v\n", list.computeChecksum())
		fmt.Printf("Time: %v\n", time2.Sub(time1))
	}
	// part 2
	{
		time1 := time.Now()
		list := &sectorList{}
		list.buildFromBytes(bytes, true)
		list.compact()
		time2 := time.Now()
		fmt.Printf("Checksum: %v\n", list.computeChecksum())
		fmt.Printf("Time: %v\n", time2.Sub(time1))
	}
}
