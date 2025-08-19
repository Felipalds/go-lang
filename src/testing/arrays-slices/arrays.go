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
