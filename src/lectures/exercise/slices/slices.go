//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

func printInfo(title string, assemblyLine []string) {
	fmt.Println()
	fmt.Println("---", title, "---")

	for i := 0; i < len(assemblyLine); i++ {
		part := assemblyLine[i]
		fmt.Println(part)
	}
}

type Part string

func main() {
	// assemblyLine := []string{"Extract", "Translate", "Load"}
	// printInfo("Initial", assemblyLine)
	// 或者使用下面的方式
	assemblyLine := make([]string, 3)
	assemblyLine[0] = "Extract"
	assemblyLine[1] = "Translate"
	assemblyLine[2] = "Load"
	printInfo("Initial", assemblyLine)

	assemblyLine = append(assemblyLine, "Query", "Create Mart")
	printInfo("Added new", assemblyLine)

	assemblyLine = assemblyLine[3:]
	printInfo("Only new", assemblyLine)

}
