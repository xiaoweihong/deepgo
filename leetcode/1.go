package main

import "fmt"

func removeDuplicates(nums []int) int {

	//var a []int
	tmp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[tmp] {
			tmp++
			nums[tmp] = nums[i]
		}
	}
	return tmp + 1

}
func removeDuplicates2(nums []int) int {
	number := 0 //记录不同数字的下标
	lens := len(nums)
	for i := 0; i < lens; i++ {
		if nums[i] != nums[number] {
			//不相同 下标前移 赋值
			number++
			nums[number] = nums[i]
		}
	}
	return number + 1
}

func main() {
	num := []int{3, 2, 1, 2}
	len := removeDuplicates(num)

	fmt.Println(len)

	for i := 0; i < len; i++ {
		print(num[i])
	}
}
