package main

import "fmt"

func main() {

	// slice
	arrF := []float32{1, 2, 3, 4}
	fmt.Printf("type: %T, len: %d, cap: %d\n", arrF, len(arrF), cap(arrF))
	fmt.Printf("res: %f\n", Sum(arrF))

	arrI := []int{1, 2, 3, 4}
	res, avg := SumAndAverage(arrI)
	fmt.Printf("res: %d, avg: %f\n", res, avg)

}

// Sum input type: slice
func Sum(arr []float32) (sum float32) {
	for i := range arr {
		sum += arr[i]
	}
	return
}

//SumAndAverage unnamed variable
func SumAndAverage(arr []int) (int, float32) {
	var sum int
	for i := range arr {
		sum += arr[i]
	}
	avg := float32(sum / len(arr))
	return sum, avg
}
