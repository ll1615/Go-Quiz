package main

import "fmt"

func main() {
	str := "234监考老师电话费35sf飞机dfd"
	fmt.Printf("%6s: %s\n", "before", str)
	fmt.Printf("%6s: %s\n", "after", string(reverse([]byte(str))))
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	for i := len(s) - 1; i >= 0; {
		n := 1
		switch s[i] & 240 {
		case 240:
			s[i], s[i-3] = s[i-3], s[i]
			s[i-1], s[i-2] = s[i-2], s[i-1]
			n = 4
		case 224:
			s[i], s[i-2] = s[i-2], s[i]
			n = 3
		case 192:
			s[i], s[i-1] = s[i-1], s[i]
			n = 2
		}
		i -= n
	}
	return s
}

// output:
// before: 234监考老师电话费35sf飞机dfd
//  after: dfd机飞fs53费话电师老考监432

// utf8:
// 0xxxxxxx                             runes 0-127    (ASCII)
// 110xxxxx 10xxxxxx                    128-2047       (values <128 unused)
// 1110xxxx 10xxxxxx 10xxxxxx           2048-65535     (values <2048 unused)
// 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx  65536-0x10ffff (other values unused)
