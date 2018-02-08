package roman_calculator

import (
	"errors"
	"fmt"
)

type OutOfBoundsError int

func (i OutOfBoundsError) Error() string {
	return fmt.Sprintf("roman_calculator.encode: %d is negative or greater then 3999", int(i))
}

func Encode(input int) (string, error) {
	if input < 0 || input > 3999 {
		return "invalid input", errors.New("Input out of bounds, 0-3999")
	}
	current = ""
	return encoding(input), nil
}

// Encode Arabic to Roman
func encoding(input int) string {
	if input == 0 {
		return "N"
	}

	for _, value := range roman {
		if input >= value.Val {
			current += value.Idx
			encoding(input - value.Val)
			break
		}
	}

	return current
}
