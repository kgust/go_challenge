package main

// TODO Write a help command and make it the default

import (
	"fmt"
	"os"
	"strconv"
	calc "vumc.org/go_challenge/roman_calculator"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(os.Args[0] + ":")
		fmt.Println("\tencode (enc) - convert to roman numerals")
		fmt.Println("\tdecode (devc) - convert from roman numerals")
		fmt.Println("\tcalculate (calc) - perform simple math operations")
		fmt.Println("")
		fmt.Println("Examples:")
		fmt.Println("\t" + os.Args[0] + " encode 1492")
		fmt.Println("\t" + os.Args[0] + " decode MCMLXX")
		fmt.Println("\t" + os.Args[0] + " calculate MCMLXX + XXXIII")
		return
	}
	command := os.Args[1]

	if command == "enc" || command == "encode" {
		input := os.Args[2]
		value, _ := strconv.Atoi(input)
		output, err := calc.Encode(value)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(output)
	}

	if command == "dec" || command == "decode" {
		input := os.Args[2]
		output, err := calc.Decode(input)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(output)
	}

	if command == "calc" || command == "calculate" {
		input1 := os.Args[2]
		operation := os.Args[3]
		input2 := os.Args[4]

		switch operation {
		case "+":
			output, err := calc.Add(input1, input2)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(output)
			}
		case "-":
			output, err := calc.Subtract(input1, input2)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(output)
			}
		case "*":
			output, err := calc.Multiply(input1, input2)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(output)
			}
		case "/":
			output, err := calc.Divide(input1, input2)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(output)
			}
		default:
			// TODO throw an error
		}
	}
}
