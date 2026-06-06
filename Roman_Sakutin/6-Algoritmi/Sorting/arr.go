package main

//func main() {
//	nums1 := []int{4, 1, 2}
//	nums2 := []int{1, 3, 4, 2}
//	fmt.Println(nextGreaterElement(nums1, nums2))
//}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var arr []int
	for i := 0; i < len(nums1); i++ {
		var j int
		for k, r := range nums2 {
			if r == nums1[i] {
				j = k
				break
			}
		}
		if j == len(nums2)-1 {
			arr = append(arr, -1)
			continue
		}
		if nums2[j] < nums2[j+1] {
			arr = append(arr, nums2[j+1])
		} else {
			arr = append(arr, -1)
		}
	}
	return arr
}
