package roman_calculator

func FindRoman(idx string, val int) RomanNumeral {
	for _, v := range roman {
		if idx == v.Idx || val == v.Val {
			return v
		}
	}
	return RomanNumeral{"N", 0}
}
