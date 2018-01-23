package roman_calculator

// TODO whitespace trim the input
// TODO string to lower input
// TODO if an invalid char is passed, throw an error
// TODO if the sum is greater than 3999, throw an error

// Decode from Roman to Arabic
func Decode(input string) int {
	Σ := 0

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

	return Σ
}
