// 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回

package main

import (
	"fmt"
	"unicode"
)

func main() {
	arr := "  123123胜多   负少  打请问覅瓯        海区沃尔  夫hi水    电费qwoeiuh2131    "
	fmt.Printf("%6s:|%s|\n", "before", arr)
	fmt.Printf("%6s:|%s|\n", "after", string(trimSpace([]byte(arr))))
}

func trimSpace(byt []byte) []byte {
	ln := len(byt)
	for i := 0; i < ln; {
		if unicode.IsSpace(rune(byt[i])) {
			j := i + 1
			for ; unicode.IsSpace(rune(byt[j])) && j < ln; j++ {
			}
			ln = ln - j + i + 1
			copy(byt[i+1:], byt[j:])
			i = j
		} else {
			i++
		}
	}
	return byt
}

// 输出:
//before:|  123123胜多   负少  打请问覅瓯        海区沃尔  夫hi水    电费qwoeiuh2131    |
// after:| 123123胜多 负少 打请问覅瓯 海区沃尔 夫hi水 电费qwoeiuh2131 |
