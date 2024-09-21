package main

import (
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) GetValue(index int) (int, error) {
	if index > len(s.values) {
		return 0, fmt.Errorf("no value at index %v", index)
	} else {
		return s.values[index], nil
	}
}

func main() {
	stuff := Stuff{}
	value, err := stuff.GetValue(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

}
