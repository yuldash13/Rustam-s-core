package main

//func main() {
//	fmt.Println(stack(")"))
//}

//func isValid(s string) bool {
//	var i, j, k int
//	for _, r := range s {
//		if r == '(' {
//			i++
//		} else if r == ')' {
//			i--
//		}
//		if r == '{' {
//			j++
//		} else if r == '}' {
//			j--
//		}
//		if r == '[' {
//			k++
//		} else if r == ']' {
//			k--
//		}
//	}
//	if i == 0 && j == 0 && k == 0 {
//		return true
//	}
//	return false
//}

func stack(s string) bool {
	var arr []int32
	var n int
	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			arr = append(arr, r)
			continue
		}
		if len(arr) == 0 {
			n = -1
			break
		}
		if arr[len(arr)-1] == '(' && r == ')' || arr[len(arr)-1] == '{' && r == '}' || arr[len(arr)-1] == '[' && r == ']' {
			arr = arr[:len(arr)-1]
			n = 0
		} else {
			break
		}
	}
	if len(arr) == 0 && n == 0 {
		return true
	}
	return false
}
