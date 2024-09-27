//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func ReadLine(lines []string) {
	sumLetters, sumDigits, sumSpaces, sumPunctuation := 0, 0, 0, 0
	for _, line := range lines {
		letters, digits, spaces, punctuation := 0, 0, 0, 0

		for _, char := range line {
			switch {
			case unicode.IsLetter(char):
				letters++
			case unicode.IsDigit(char):
				digits++
			case unicode.IsSpace(char):
				spaces++
			case unicode.IsPunct(char):
				punctuation++
			}
		}
		// Print the report for the current line
		fmt.Printf("Line: %s\n", line)
		fmt.Printf("Letters: %d\n", letters)
		fmt.Printf("Digits: %d\n", digits)
		fmt.Printf("Spaces: %d\n", spaces)
		fmt.Printf("Punctuation: %d\n\n", punctuation)
		sumLetters += letters
		sumDigits += digits
		sumSpaces += spaces
		sumPunctuation += punctuation

	}
	fmt.Printf("Sum Letters: %d\n", sumLetters)
	fmt.Printf("Sum Digits: %d\n", sumDigits)
	fmt.Printf("Sum Spaces: %d\n", sumSpaces)
	fmt.Printf("Sum Punctuation: %d\n\n", sumPunctuation)
}

func processText(lines []string, op func(lines []string)) {
	// return op(lines) // 因为这里没有return任何东西
	op(lines)
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	processText(lines, ReadLine)
}
