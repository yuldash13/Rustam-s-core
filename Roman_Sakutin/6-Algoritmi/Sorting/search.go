package main

import (
	"math/rand"
)

//func main() {
//	arr := []int{2, 5}
//	fmt.Println(search(arr, 2))
//}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	index := rand.Intn(right + 1)
	q := nums[index]
	i := left
	j := right
	for i <= j {
		for nums[i] < q {
			i++
		}
		for q < nums[j] {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	quickSort(nums, i, right)
	quickSort(nums, left, j)
}

//func binarySearch(nums []int, target int) int {
//	left := -1
//	right := len(nums)
//	for right-left > 1 {
//		mid := (right + left) / 2
//		if nums[mid] >= target {
//			right = mid
//		} else {
//			left = mid
//		}
//	}
//	return right
//}

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := (r + l) / 2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid
		}
	}
	if nums[l] != target && nums[r] != target {
		r = -1
	}
	return r
}
