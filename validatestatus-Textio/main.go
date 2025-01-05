package main

import (
	"errors"
)

func validateStatus(status string) error {
	//empty status?
	if status == "" || len(status) == 0 {
		return errors.New("status cannot be empty")
	}

	//status too long?
	if len(status) > 140 {
		return errors.New("status exceeds 140 characters")
	}

	return nil
}
