package main

func Sum(x int, y int) int {
	return x + y + 1 //will fail because of 1
}

func main() {
	Sum(5, 5)
}
