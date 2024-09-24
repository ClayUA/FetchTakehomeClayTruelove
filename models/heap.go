package models

import "container/heap"

// creating a heap that can store and pop the oldest transactions
// used these docs for reference https://pkg.go.dev/container/heap#Pop
type TransactionHeap []Transaction

// length is required to access the minimum item
func (heap TransactionHeap) Len() int {
	return len(heap)
}

// compare timestamps to create a min heap of timestamps
func (heap TransactionHeap) Less(i, j int) bool {
	return heap[i].ParsedTime.Before(heap[j].ParsedTime)
}

// swaps the location of items
func (heap TransactionHeap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap *TransactionHeap) Push(item interface{}) {
	*heap = append(*heap, item.(Transaction))
}

// removes oldest transaction from heap
func (heapPointer *TransactionHeap) Pop() interface{} {
	heap := *heapPointer
	lengthOfHeap := len(heap)
	oldestTransaction := heap[lengthOfHeap-1]
	*heapPointer = heap[0 : lengthOfHeap-1]
	return oldestTransaction
}

// pushes a transaction onto the heap
func (heapPointer *TransactionHeap) AddTransaction(currentTransaction Transaction) {
	heap.Push(heapPointer, currentTransaction)
}

// pops the oldest transaction
func (heapPointer *TransactionHeap) RemoveOldestTransaction() Transaction {
	return heap.Pop(heapPointer).(Transaction)
}

// accesses the oldest transaction so we can compare before popping
func (heapPointer *TransactionHeap) CheckOldestTransaction() Transaction {
	return (*heapPointer)[0]
}

// if oldest transaction has more points than required we need to update it without removing it
// nothing is changed internally in the heap because timestamp remains the same
func (heapPointer *TransactionHeap) UpdateOldestTransaction(pointsToDeduct int) {
	oldestTransaction := &(*heapPointer)[0]
	oldestTransaction.Points -= pointsToDeduct
	return
}

// creating new heap
func InitializeHeap() *TransactionHeap {
	newHeap := &TransactionHeap{}
	heap.Init(newHeap)
	return newHeap
}
