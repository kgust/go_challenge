package roman_calculator

// TODO if the input is negative, throw an error
// TODO if the input > 3999, throw an error

func Encode(input int) string {
	current = ""
	return encoding(input)
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
