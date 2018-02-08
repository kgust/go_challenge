package roman_calculator

import (
	"errors"
	"fmt"
	"strings"
)

// Decode from Roman to Arabic
func Decode(input string) (int, error) {
	Σ := 0

	input = strings.ToUpper(strings.TrimSpace(input))

	if input == "N" {
		return 0, nil
	}

	for {
		α := ""
		β := ""
		length := len(input)

		if len(input) == 0 {
			break
		} else if len(input) == 1 {
			β = input[0:1]
		} else {
			α = input[0:2]
			β = input[0:1]
		}

		for _, v := range roman {
			if α == v.Idx {
				Σ += v.Val
				input = input[2:]
				break
			} else if β == v.Idx {
				Σ += v.Val
				input = input[1:]
				break
			}
		}

		if length == len(input) {
			// didn't change, loop detected
			err := errors.New(fmt.Sprintf("decode: invalid characters detected (%s)", α))
			return -1, err
		}
	}

	if Σ > 3999 {
		return -1, errors.New("Incorrect input: roman numerals greater than 3999 are not allowed")
	}

	return Σ, nil
}
