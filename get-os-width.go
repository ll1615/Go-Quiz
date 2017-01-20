// print the byte-wide of current using Operating system (当前系统位宽)
package main

import (
	"fmt"
)

func main(){
	fmt.Println(32 << (^uint(0) >> 63))
}
