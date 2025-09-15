package main

type STree struct {
	arr []int
}

func NewSTree(arr []int) *STree {
	size := 1
	for size < len(arr) {
		size *= 2
	}
	arr1 := make([]int, 2*size)
	return &STree{arr: arr1}
}

func (st *STree) query(l, r int) int {
	if l/2 == r/2 {
		return st.arr[l/2]
	} else {
		if l%2 != 0 {
			return st.arr[l] + st.query(l+1, r)
		}
		if r%2 != 1 {
			return st.arr[r] + st.query(l, r-1)
		}
	}
	return st.query(l/2, r/2)
}

func (st *STree) build(arr []int) {
	for i := 0; i < len(arr); i++ {
		st.arr[(len(st.arr)/2)+i] = arr[i]
	}
	for i := (len(st.arr) / 2) - 1; i > 0; i-- {
		st.arr[i] = st.arr[i*2] + st.arr[i*2+1]
	}
}

func (st *STree) update(i, value int) {
	st.arr[i] = value
	for j := i / 2; j > 0; j /= 2 {
		st.arr[j] = st.arr[j*2] + st.arr[j*2+1]
	}
}
