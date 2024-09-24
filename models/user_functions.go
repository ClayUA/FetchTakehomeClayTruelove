package models

// Adds points to a users account based on payer
func (user *User) AddUserPoints(t *Transaction) error {
	user.Mutex.Lock()
	defer user.Mutex.Unlock()

	user.TotalPoints += t.Points
	user.PayerMap[t.Payer] += t.Points
	user.SortedTransactions.AddTransaction(*t)
	return nil

}

// Accesses and returns map of all payer balances from a user
func (user *User) RetrieveUserBalance() map[string]int {
	user.Mutex.Lock()
	defer user.Mutex.Unlock()

	return CurrentUser.PayerMap
}

// spends a users points by accessing the oldest transaction using a min heap
// pop used transactions and update transactions that contain more points than PointsToRemove
// keep track of each payers spent points using a map and return it
func (user *User) SpendUserPoints(PointsToRemove int) map[string]int {
	user.Mutex.Lock()
	defer user.Mutex.Unlock()

	spentPointsMap := make(map[string]int)

	for PointsToRemove > 0 {

		oldestTransaction := user.SortedTransactions.CheckOldestTransaction()

		if oldestTransaction.Points <= PointsToRemove {
			poppedTransaction := user.SortedTransactions.RemoveOldestTransaction()
			PointsToRemove -= poppedTransaction.Points
			spentPointsMap[poppedTransaction.Payer] -= poppedTransaction.Points
			user.PayerMap[poppedTransaction.Payer] -= poppedTransaction.Points
			user.TotalPoints -= poppedTransaction.Points
		} else {
			user.SortedTransactions.UpdateOldestTransaction(PointsToRemove)
			spentPointsMap[oldestTransaction.Payer] -= PointsToRemove
			user.PayerMap[oldestTransaction.Payer] -= PointsToRemove
			user.TotalPoints -= PointsToRemove
			PointsToRemove = 0
		}
	}
	return spentPointsMap
}
