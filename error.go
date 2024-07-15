package main

type ZeroDivisionError struct{}

func (e *ZeroDivisionError) Error() string {
	return "division by zero"
}
