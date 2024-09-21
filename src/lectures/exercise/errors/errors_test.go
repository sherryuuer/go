package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"19:20:13", true},
		{"9:26:1", true},
		{"19:20:-1", false},
		{"19:20:Y", false},
		{"19:20", false},
		{"E:20:13", false},
		{"", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("Error should be nil: time(%v)/err(%v)", data.time, err)
		}
		// 原本的课程中并没有这段false对象的检查，于是我没有检查出错误的代码
		// 对second部分的判断我漏改了，但是加上这段后我测试出了这个问题
		// 测试样例全覆盖很重要
		if !data.ok && err == nil {
			t.Errorf("Error expected but got nil: time(%v)", data.time)
		}
	}
}
