package model

type TestResult struct {
	Name     string
	Passed   bool
	Errors   []string
	Duration int64
}
