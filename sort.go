// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binheap

type intType int

func (i intType) Priority() int {
    return int(i)
}

// heap-sort int slice
func SortInt(slice []int) {

    heap := NewSize(len(slice) + 2)

    for i := 0; i < len(slice); i++ {
        heap.Add(intType(slice[i]))
    }

    for i := 0; i < len(slice); i++ {
        slice[i] = heap.Remove().Priority()
    }
}

// heap-sort Heapable slice
func Sort(slice []Heapable) {

    heap := NewSize(len(slice) + 2)

    for i := 0; i < len(slice); i++ {
        heap.Add(slice[i])
    }

    for i := 0; i < len(slice); i++ {
        slice[i] = heap.Remove()
    }
}
