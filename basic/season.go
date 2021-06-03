package main

import "fmt"

func main() {
	fmt.Println(season(7))
	fmt.Println(season(12))
	fmt.Println(season(9))
	fmt.Println(season(100))
}

// 输入月份, 返回 season
func season(month int) string {

	switch month {
	case 12, 1, 2:
		return "winter"
	case 3, 4, 5:
		return "spring"
	case 6, 7, 8:
		return "summer"
	case 9, 10, 11:
		return "autumn"
	default:
		return "wrong para!"
	}


}
