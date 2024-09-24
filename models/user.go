package models

import (
	"sync"
)

// only using one user for this exercise
type User struct {
	UserID             int
	PayerMap           map[string]int
	SortedTransactions *TransactionHeap
	TotalPoints        int
	mu                 sync.Mutex
}

// global user instance we will use for this assessment
var CurrentUser = User{
	UserID:             1,
	PayerMap:           make(map[string]int),
	TotalPoints:        0,
	SortedTransactions: InitializeHeap(),
}
