package main

import "errors"

func divide(x, y float64) (float64, error) {
	if y == 0 {
		// no divinding by 0
		err := errors.New("no dividing by 0")
		return 0.00, err
	}
	return x / y, nil
}
