package main
import "fmt"


func main() {
	//Inputs
	A := []int{1,3,1,4,2,3,5,4}
	x := 5
	// A := []int{1,2,3,4,3,5}
	// x := 5
	//fmt.Println(Solution(x,A))
	fmt.Println(FrogRiverOne(x,A))
}
//solutionOnlineThatPassed (despite not appropriately handling farther out leaves)
func FrogRiverOne(X int, A []int) int {
	arraySum := 0
	seen := make(map[int]bool)

	expectedSum := X * (X + 1) / 2

	for index, value := range A {
		if !seen[value] {
			arraySum += value
			seen[value] = true
		}

		if arraySum == expectedSum {
			return index
		}
	}

	return -1
}



func Solution(X int, A []int) int {
	//make slice and a step sum
	var n []int
	p := 0

	//determine required positions and place in array
	for i := X; i >= 0; i-- {
		//add each step to slice (Start6,5,4,3,2,1)
		n = append(n,i)
	}

	// for every ascending element in array A starting at 0
	for s:= 0 ; s < len(A); s++ {
		//if all elements have not been removed
		if len(n) > 1 {
			//if that leaf is still found in n
			_ , found := Find(n, A[s])
			if found {
				//remove the leaf from leaf list
				n = remove(n, A[s])
				fmt.Println("Leaf extracted")
				//add 1 to solution
				p +=1
			}
			//if leaf not in the list(or already removed)
			fmt.Println("Leaf not found or redundant")
			continue
		}
		//if steps-list has reached zero
		continue
	}
	return p
}


func remove(slice []int, s int) []int {
	return append(slice[:s+1], slice[s:]...)
}

//order unimportant
//func remove(s []int, i int) []int {
//	s[len(s)-1], s[i] = s[i], s[len(s)-1]
//	return s[:len(s)-1]
//}

func Find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
