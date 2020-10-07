package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

//append only allocate new arrays, which is passed by value and will be lost once out of scope
func double(slice []int) []int {
	slice = append(slice, slice...)
	return slice
}

func mapSlice(f func(a int) int, slice []int) {
	for i, num := range slice {
		slice[i] = f(num)
	}
}

func mapArray(f func(a int) int, array [3]int) [3]int {
	var newArray [3]int
	for i := 0; i < 3; i++ {
		newArray[i] = f(array[i])
	}
	return newArray
}

func main() {
	intsSlice := []int{1, 2, 3, 4, 5}
	newSlice := intsSlice[1:3]
	mapSlice(addOne, intsSlice)
	// mapSlice(square, newSlice)
	intsSlice = double(intsSlice)
	// fmt.Println(cap(intsSlice))
	// intsSlice = append(intsSlice, intsSlice...)
	fmt.Println(newSlice)
	fmt.Println(intsSlice)
	intsArray := [3]int{1, 2, 3}
	intsArray = mapArray(addOne, intsArray)
	fmt.Println(intsArray)
}
