// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binheap_test


import(
    "rand"
    "sort"
    "testing"
    . "github.com/bjarneh/binheap"
)

type heaper struct {
    i int
}

func (h *heaper) Priority() int {
    return h.i
}


func randomInts(size int) (slice []int) {

    slice = make([]int, size)

    for i := 0; i < size; i++ {
        slice[i] = rand.Int()
    }

    return slice
}

func randomHeapable(size int) (slice []Heapable) {

    ints := randomInts(size)

    slice = make([]Heapable, size)

    for i := 0; i < size; i++ {
        slice[i] = &heaper{ints[i]}
    }

    return slice
}


func TestIntSort(t *testing.T) {

    nums := randomInts(1000)

    SortInt(nums)

    for i := 1; i < len(nums); i++ {
        if nums[i-1] > nums[i] {
            t.Fatalf("binheap.SortInts() does not sort\n")
        }
    }
}

func TestSort(t *testing.T) {

    nums := randomHeapable(1000)

    Sort(nums)

    for i := 1; i < len(nums); i++ {
        if nums[i-1].Priority() > nums[i].Priority() {
            t.Fatalf("binheap.Sort() does not sort\n")
        }
    }
}

func TestHeap(t *testing.T) {

    heap := NewSize(1) // make it short to test resize
    nums := randomInts(1000)
    elmts := make([]Heapable, len(nums))

    for i := 0; i < len(nums); i++ {
        elmts[i] = &heaper{nums[i]}
    }

    heap.AddSlice(elmts)

    if heap.Len() != len(nums) {
        t.Fatalf("heap.Len() != %d\n", len(nums))
    }

    sort.SortInts(nums)

    for i := 0; i < len(nums); i++ {
        if heap.Remove().Priority() != nums[i] {
            t.Fatalf("binheap: Priority() != %d\n", nums[i])
        }
    }

    if ! heap.Empty() {
        t.Fatalf("heap not empty\n")
    }

    nums = randomInts(1000)

    // add single elements.. i.e. not slice
    for i := 0; i < len(nums); i++ {
        heap.Add(&heaper{nums[i]})
    }

    sort.SortInts(nums)

    for i := 0; i < len(nums); i++ {
        if heap.Remove().Priority() != nums[i] {
            t.Fatalf("binheap: Priority() != %d\n", nums[i])
        }
    }

}

func BenchmarkStdlibSortInts(b *testing.B) {
    for i := 0; i < b.N; i++ {
        nums := randomInts(1000)
        sort.SortInts(nums)
    }
}

func BenchmarkHeapSortInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        nums := randomInts(1000)
        SortInt(nums)
    }
}
