package main

import "errors"

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch plan {
	case planPro:
		return messages[0:], nil
	case planFree:
		return messages[0:2], nil
	default:
		return []string{}, errors.New("unsupported plan")
	}

}
