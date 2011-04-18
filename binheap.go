// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binheap

/*
    A small binary-heap, or priority queue.
    To add an elements to heap it must implement 
    the Heapable interface:

    <pre>
    type Heapable struct {
        Priority() int
    }
    </pre>

    Typical use:

    <pre>

    type Job int

    func (j Job) Priority() int {
        return int(j)
    }

    heap := binheap.New()

    heap.Add(Job(4))
    heap.Add(Job(1))
    heap.Add(Job(10))
    .
    .

    mostUrgent := heap.Remove() // Job(1)

    // heap-sort, 1/2 speed of sort.Sort (quick-sort)

    binheap.Sort([]Job)

    </pre>

*/

import (
    "fmt"
)

type Heapable interface {
    Priority() int
}

type Heap struct {
    max  int
    tree []Heapable
}

// new Heap with default size: 50
func New() *Heap {
    return &Heap{0, make([]Heapable, 50)}
}

func NewSize(size int) *Heap {
    return &Heap{0, make([]Heapable, size)}
}

func (h *Heap) String() string {
    var res string
    for i := 1; i <= h.max; i++ {
        res = fmt.Sprintf("%s %d", res, h.tree[i].Priority())
    }
    return fmt.Sprintf("[%s ]",res)
}

func (h *Heap) Add(v Heapable) {
    if h.max+1 >= len(h.tree) {
        h.resize()
    }
    h.max++
    h.tree[h.max] = v
    h.up(h.max)
}

func (h *Heap) AddSlice(s []Heapable) {
    for i := 0; i < len(s); i++ {
        h.Add( s[i] )
    }
}

func (h *Heap) Remove() Heapable {

    if h.max == 0 {
        return nil
    }

    min := h.tree[1] // root
    h.tree[1] = h.tree[h.max]
    h.tree[h.max] = nil
    h.max--
    h.down(1)

    return min
}

func (h *Heap) Empty() bool {
    return h.max == 0
}

func (h *Heap) Len() int {
    return h.max
}

// index of son with highest pri, or -1 if no children
func (h *Heap) highPrioritySon(first, second int) int {

    // they are both larger than last element in heap
    if first > h.max || h.tree[first] == nil {
        return -1
    } else if h.tree[second] == nil {
        return first
    } else if h.tree[first].Priority() > h.tree[second].Priority() {
        return second
    }

    return first
}

func (h *Heap) up(pos int) {

    // you are root
    if pos == 1 {
        return
    }

    // less is more (son < father)
    if h.tree[pos].Priority() < h.tree[pos/2].Priority() {
        h.swap(pos/2, pos)
        h.up(pos / 2)
    }
}

func (h *Heap) down(father int) {

    son := h.highPrioritySon(father*2, (father*2)+1)

    if son != -1 {
        if h.tree[father].Priority() > h.tree[son].Priority() {
            h.swap(father, son)
            h.down(son)
        }
    }
}

func (h *Heap) swap(i, j int) {
    h.tree[j], h.tree[i] = h.tree[i], h.tree[j]
}

// double capacity
func (h *Heap) resize() {
    h.tree = append(h.tree, make([]Heapable, len(h.tree))...)
}
