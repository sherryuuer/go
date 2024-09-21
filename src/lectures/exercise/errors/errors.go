//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minite, second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

// * Example time string: 14:07:33
func ParseTime(input string) (Time, error) {
	components := strings.Split(input, ":")
	// 检查分割后长度是否正确
	if len(components) != 3 {
		return Time{}, &TimeParseError{"Invaild value", input}
	} else {
		// 检查每个部分的转换（字符 to 数值）是否无错误
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour: %v", err), input}
		}
		minite, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing minite: %v", err), input}
		}
		second, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing second: %v", err), input}
		}

		// 检查每个位置的数字是否合法
		if hour < 0 || hour > 23 {
			return Time{}, &TimeParseError{"hour out of range (0, 23)", fmt.Sprintf("%v", hour)}
		}
		if minite < 0 || minite > 59 {
			return Time{}, &TimeParseError{"minite out of range (0, 59)", fmt.Sprintf("%v", minite)}
		}
		if second < 0 || second > 59 {
			return Time{}, &TimeParseError{"second out of range (0, 23)", fmt.Sprintf("%v", second)}
		}

		// 返回最终结果
		return Time{hour, minite, second}, nil
	}
}
