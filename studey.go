package main

import "fmt"

func main() {
	//冒泡排序
	nums := [...]int{5, 1, 2, 8, 9, 3, 10}

	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Println("排序结果：", nums)
}
