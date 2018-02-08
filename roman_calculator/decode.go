package roman_calculator

import "strings"

// TODO handle unexpected input (e.g. "N", garbage chars, etc.)
// TODO if the sum is greater than 3999, throw an error

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
	}

	return Σ, nil
}
