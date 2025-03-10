package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {

	origMsgs := [3]string{primary, secondary, tertiary}

	var cost [3]int
	totCost := 0
	for i, msg := range origMsgs {
		totCost += len(msg)
		cost[i] = totCost
	}

	return origMsgs, cost
}
