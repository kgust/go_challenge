package roman_calculator

// Validation:
// https://www.rapidtables.com/convert/number/roman-numerals-converter.html

import (
	"fmt"
	"testing"
)

func TestFindRoman(t *testing.T) {
	values := map[string]int{
		"N":  0,
		"M":  1000,
		"CM": 900,
		"CD": 400,
		"X":  10,
		"IV": 4,
		"I":  1,
	}

	for idx, val := range values {
		actual := FindRoman(idx, 0)
		if actual.Val != val {
			t.Errorf("FindRoman(%d): expected %s, actual %s", idx, val, actual.Val)
		}
	}

	for idx, val := range values {
		actual := FindRoman("", val)
		if actual.Idx != idx {
			t.Errorf("FindRoman(%d): expected %s, actual %s", val, idx, actual.Idx)
		}
	}
}

func TestEncode(t *testing.T) {
	values := map[int]string{
		0:    "N",
		1:    "I",
		2:    "II",
		3:    "III",
		4:    "IV",
		5:    "V",
		6:    "VI",
		7:    "VII",
		8:    "VIII",
		9:    "IX",
		10:   "X",
		11:   "XI",
		1967: "MCMLXVII",
		1999: "MCMXCIX",
		2000: "MM",
		2018: "MMXVIII",
		3999: "MMMCMXCIX",
	}
	for input, expected := range values {
		actual, err := Encode(input)

		if err != nil {
			fmt.Println(err)
		}

		if actual != expected {
			t.Errorf("Encode(%d): expected %s, actual %s", input, expected, actual)
		}
	}
}

func TestDecode(t *testing.T) {
	values := map[string]int{
		"I":         1,
		"II":        2,
		"III":       3,
		"IV":        4,
		"V":         5,
		"VI":        6,
		"VII":       7,
		"VIII":      8,
		"IX":        9,
		"X":         10,
		"XI":        11,
		"XII":       12,
		"XIII":      13,
		"XIV":       14,
		"XV":        15,
		"XVI":       16,
		"XVII":      17,
		"XVIII":     18,
		"XIX":       19,
		"XX":        20,
		"XXI":       21,
		"XXII":      22,
		"XXIII":     23,
		"XXIV":      24,
		"XXV":       25,
		"MCMLXVII":  1967,
		"MCMXCIX":   1999,
		"MM":        2000,
		"MMXVIII":   2018,
		"MMMCMXCIX": 3999,
	}

	for input, expected := range values {
		actual, err := Decode(input)

		if err != nil {
			fmt.Println(err)
		}

		if actual != expected {
			t.Errorf("Decode(%d): expected %d, actual %d", input, expected, actual)
		}
	}
}
