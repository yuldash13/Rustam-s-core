package main

//func main() {
//	nums := []int{10, 2, 2, 0, 0, 4}
//	quickSort(nums, 0, len(nums)-1)
//	fmt.Println(nums)
//}
//
//func quickSort(nums []int, left, right int) {
//	if left >= right {
//		return
//	}
//	i := left
//	j := right
//	index := (left + right) / 2
//	q := nums[index]
//	for i <= j {
//		for nums[i] < q {
//			i++
//		}
//		for q < nums[j] {
//			j--
//		}
//		if i <= j {
//			nums[i], nums[j] = nums[j], nums[i]
//			i++
//			j--
//		}
//	}
//	quickSort(nums, i, right)
//	quickSort(nums, left, j)
//}
