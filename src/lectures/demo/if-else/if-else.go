package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	student1Score, student2Score, student3Score := 9, 8, 7

	if student1Score > student2Score {
		fmt.Println("Student1 got a higher score")
	} else if student1Score < student2Score {
		fmt.Println("Student2 got a higher score")
	} else {
		fmt.Println("Student1 and student2 got the same score")
	}

	if average(student1Score, student2Score, student3Score) > 7 {
		fmt.Println("Good Job")
	} else {
		fmt.Println("Instructor did a bad job")
	}

}
