package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 1, 5, 6, 7}
	fmt.Println(sort31(arr))

	arr1 := []int{1, 7, 7, 15}
	arr2 := []int{2, 6, 8, 13}
	fmt.Println(merge32(arr1, arr2))

	arr3a := []int{3, 2, 2, 1}
	fmt.Println(selectionSort33a(arr3a))

	arr3b := []int{3, 2, 2, 1}
	fmt.Println(insertionSort33b(arr3b))

	//34. Число обменов равно длине массива минус один. Да.

	arr11 := []int{1, 2, 7, 15}
	arr22 := []int{2, 6, 7, 13}
	fmt.Println(findEqual35(arr11, arr22))

	arr111 := []int{1, 2, 7, 15}
	arr222 := []int{2, 6, 7, 13}
	fmt.Println(findMin36(arr111, arr222))

	arr1111 := []int{1, 2, 7, 15}
	arr2222 := []int{2, 6, 7, 14}
	s := 14
	fmt.Println(findSum37(arr1111, arr2222, s))

	arr11111 := []int{1, 2, 7, 15}
	arr22222 := []int{2, 6, 7, 14}
	fmt.Println(findQuan38(arr11111, arr22222))

	arr111111 := []int{1, 2, 7, 15}
	arr222222 := []int{2, 6, 7, 14}
	fmt.Println(findBig39(arr111111, arr222222))
}

func sort31(arr []int) []int {
	var n int
	for i := 0; n < len(arr)-1; i++ {
		if i == len(arr)-1 {
			i = 0
			n++
		}
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	return arr
}

func merge32(arr1, arr2 []int) []int {
	var arr []int
	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr = append(arr, arr1[i])
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
		}
	}
	return arr
}

func selectionSort33a(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		mini := i
		for j := i + 1; j < len(arr); j++ {
			if arr[mini] > arr[j] {
				mini = j
			}
		}
		arr[mini], arr[i] = arr[i], arr[mini]
	}
	return arr
}

func insertionSort33b(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		j := i
		if j > 0 && arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}
	}
	return arr
}

func findEqual35(arr1, arr2 []int) int {
	lens := max(len(arr1), len(arr2))
	n := 0
	j := 0
	for i := 0; ; i++ {
		if i == lens {
			j++
			i = 0
		}
		if j == lens {
			break
		}
		if arr1[j] == arr2[i] {
			n++
		}
	}
	return n
}

func findMin36(arr1, arr2 []int) (int, int) {
	lens := max(len(arr1), len(arr2))
	mini := 0
	mini1 := 0
	j := 0
	a := 0
	b := 0
	for i := 0; ; i++ {
		if i == lens {
			i = 0
			j++
		}
		if j == lens {
			break
		}
		if arr1[i] < arr2[j] {
			mini1 = arr2[j] - arr1[i]
		} else if arr1[i] > arr2[j] {
			mini1 = arr1[i] - arr2[j]
		}
		if mini == 0 {
			mini = mini1
		}
		if mini > mini1 {
			mini = mini1
			a = i
			b = j
		}
	}
	return a, b
}

func findSum37(arr1, arr2 []int, s int) (int, int) {
	j := 0
	i := 0
	lens := max(len(arr1), len(arr2))
	for ; ; i++ {
		if i == lens {
			j++
			i = 0
		}
		if j == lens {
			break
		}
		if arr1[j]+arr2[i] == s {
			break
		}
	}
	return j, i
}

func findQuan38(arr1, arr2 []int) int {
	var quan, j int
	lens := max(len(arr1), len(arr2))
	for i := 0; ; i++ {
		if i == lens {
			i = 0
			j++
		}
		if j == lens {
			break
		}
		if arr1[i] == arr2[j] {
			quan++
		}
	}
	return quan
}

func findBig39(arr1, arr2 []int) int {
	var big, j int
	lens := max(len(arr1), len(arr2))
	for i := 0; ; i++ {
		if i == lens {
			i = 0
			j++
		}
		if j == lens {
			break
		}
		if arr1[i] > arr2[j] {
			big++
		}
	}
	return big
}
