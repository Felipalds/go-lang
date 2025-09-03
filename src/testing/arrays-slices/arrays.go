package arraysslices

func SumArray(arr [5]int) (sum int) {

	for _, v := range arr {
		sum += v
	}
	return
}

func SumSlice(sli []int) (sum int) {
	for _, v := range sli {
		sum += v
	}
	return
}

func SumAllSmart(numbers ...[]int) []int {
	length := len(numbers)
	resp := make([]int, length)

	for i, slc := range numbers {
		resp[i] = SumSlice(slc)
	}
	return resp
}

func SumAllDumb(numbers ...[]int) []int {
	var resp []int
	for _, slc := range numbers {
		resp = append(resp, SumSlice(slc))
	}
	return resp
}
