package models

func (user *User) AddUserPoints(t *Transaction) error {
	user.mu.Lock()
	defer user.mu.Unlock()

	user.TotalPoints += t.Points
	user.PayerMap[t.Payer] += t.Points
	user.SortedTransactions.AddTransaction(*t)
	return nil

}

func (user *User) RetrieveUserBalance() map[string]int {
	user.mu.Lock()
	defer user.mu.Unlock()

	return CurrentUser.PayerMap
}

func (user *User) SpendUserPoints(PointsToRemove int) map[string]int {
	user.mu.Lock()
	defer user.mu.Unlock()

	SpentPointsMap := make(map[string]int)

	for PointsToRemove > 0 {

		oldestTransaction := user.SortedTransactions.CheckOldestTransaction()

		if oldestTransaction.Points <= PointsToRemove {
			poppedTransaction := user.SortedTransactions.RemoveOldestTransaction()
			PointsToRemove -= poppedTransaction.Points
			SpentPointsMap[poppedTransaction.Payer] -= poppedTransaction.Points
			user.PayerMap[poppedTransaction.Payer] -= poppedTransaction.Points
			user.TotalPoints -= poppedTransaction.Points
		} else {
			user.SortedTransactions.UpdateOldestTransaction(PointsToRemove)
			SpentPointsMap[oldestTransaction.Payer] -= PointsToRemove
			user.PayerMap[oldestTransaction.Payer] -= PointsToRemove
			user.TotalPoints -= PointsToRemove
			PointsToRemove = 0
		}
	}
	return SpentPointsMap
}
