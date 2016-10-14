// 找出缺失的整数:一个无序数组里有若干个正整数，范围从1到100，其中98个整数都出现了偶数次，只有 两个 整数出现了奇数次(比如1,1,2,2,3,4,5,5)
// 如何找到这个出现奇数次的整数？

package main

import "fmt"

func main() {

	nums := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13,
		13, 14, 14, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 23, 23, 24, 24, 25,
		25, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31, 31, 32, 78, 33, 33, 34, 34, 35, 35, 36, 36, 37,
		37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 45, 45, 46, 46, 47, 47, 48, 48, 49, 49, 50, 50}

	fmt.Println(findMissingNum(nums))
}

func findMissingNum(nums []int) (int, int) {
	ln := len(nums)

	var xor int
	for i := 0; i < ln; i++ {
		xor ^= nums[i]
	}

	diffIndex := 1
	for xor&1 != 1 {
		diffIndex = diffIndex << 1
		xor = xor >> 1
	}

	var n1, n2 int
	for i := 0; i < ln; i++ {
		if diffIndex&nums[i] == diffIndex {
			n1 ^= nums[i]
		} else {
			n2 ^= nums[i]
		}
	}

	return n1, n2
}

// output:78 32
