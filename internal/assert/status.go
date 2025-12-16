package assert

import "fmt"

func Status(actual, expected int) error {
	if actual != expected {
		return fmt.Errorf("status mismatch: expected %d, got %d", expected, actual)
	}
	return nil
}
