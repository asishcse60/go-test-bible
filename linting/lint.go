package linting

import "fmt"

func checkFlag(flag bool) bool {
	return flag
}

func errChecking() (int, error) {
	a := 2
	fmt.Println(a)
	return a, nil
}

func callsErrChecking() int {
	val, _ := errChecking()
	return val
}
