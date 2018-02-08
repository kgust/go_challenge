package roman_calculator

import "errors"

func Add(input1, input2 string) (string, error) {
	value1, err := Decode(input1)
	if err != nil {
		return "An error occurred", err
	}

	value2, err := Decode(input2)
	if err != nil {
		return "An error occurred", err
	}

	return Encode(value1 + value2)
}

func Subtract(input1, input2 string) (string, error) {
	value1, err := Decode(input1)
	if err != nil {
		return "An error occurred", err
	}

	value2, err := Decode(input2)
	if err != nil {
		return "An error occurred", err
	}

	if value2 > value1 {
		return "negative values are not allowed", errors.New("Negative values are not allowed")
	}

	return Encode(value1 - value2)
}

func Multiply(input1, input2 string) (string, error) {
	value1, err := Decode(input1)
	if err != nil {
		return "An error occurred", err
	}

	value2, err := Decode(input2)
	if err != nil {
		return "An error occurred", err
	}

	return Encode(value1 * value2)
}

func Divide(input1, input2 string) (string, error) {
	value1, err := Decode(input1)
	if err != nil {
		return "An error occurred", err
	}

	value2, err := Decode(input2)
	if err != nil {
		return "An error occurred", err
	}
	if value2 == 0 {
		return "Divide by zero error", errors.New("Divide by zero error")
	}

	remainder := value1 % value2

	if remainder != 0 {
		result, err := Encode(value1 / value2)
		if err != nil {
			return "A error occurred encoding result", err
		}

		remainder, err := Encode(remainder)
		if err != nil {
			return "A error occurred encoding remainder", err
		}

		return result + " (remainder: " + remainder + ")", nil
	}

	return Encode(value1 / value2)
}
