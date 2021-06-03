package main

var num int = 10
var num1, num2 int

// Go 没有重写和继承, 为了提高效率
func main() {

}

// unnamed return variables
func getX2AndX3(input int) (int, int) {
	return input * 2, input * 3
}

// named return variables
// recommended, clearer, shorter, self-documenting
func getX2AndX3_2(input int) (x2, x3 int) {
	x2 = input * 2
	x3 = input * 3
	// return x2, x3
	return
}
