package models

import "container/heap"

// creating a heap that can store and pop the oldest transactions
type TransactionHeap []Transaction

// lenght is required to access the minimum item
func (h TransactionHeap) Len() int {
	return len(h)
}

// compare timestamps to create a min heap of timestamps
func (h TransactionHeap) Less(i, j int) bool {
	return h[i].ParsedTime.Before(h[j].ParsedTime)
}

// swaps the location of items
func (h TransactionHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *TransactionHeap) Push(x interface{}) {
	*h = append(*h, x.(Transaction))
}

// removes oldest transaction from heap
func (h *TransactionHeap) Pop() interface{} {
	heap := *h
	lengthOfHeap := len(heap)
	oldestTransaction := heap[lengthOfHeap-1]
	*h = heap[0 : lengthOfHeap-1]
	return oldestTransaction
}

// pushes a transaction onto the heap
func (h *TransactionHeap) AddTransaction(t Transaction) {
	heap.Push(h, t)
}

// pops the oldest transaction
func (h *TransactionHeap) RemoveOldestTransaction() Transaction {
	return heap.Pop(h).(Transaction)
}

// accesses the oldest transaction so we can compare before popping
func (h *TransactionHeap) CheckOldestTransaction() Transaction {
	return (*h)[0]
}

// if oldest transaction has more points than required we need to update it without removing it
// nothing is changed internally in the heap because timestamp remains the same
func (h *TransactionHeap) UpdateOldestTransaction(pointsToDeduct int) {
	oldestTransaction := &(*h)[0]
	oldestTransaction.Points -= pointsToDeduct
	return
}

func InitializeHeap() *TransactionHeap {
	h := &TransactionHeap{}
	heap.Init(h)
	return h
}
