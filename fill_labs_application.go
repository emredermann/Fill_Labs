package main

import(
	"fmt"
	"strings"
	"sort"
	)
	/*
	Q1) Write a function that sorts a bunch of words by the number of character “a”s within the
		word (decreasing order). If some words contain the same amount of character “a”s then you
		need to sort those words by their lengths.
	*/
	/*
	Q2) Write a recursive function which takes one integer parameter. Please bear in mind that
		finding the algorithm needed to generate the output below is the main point of the question.
	*/
	/*
	Q3) Write a function which takes one parameter as an array/list. Find most repeated data
		within a given array.
	*/

func q_1(){
	var theArray  [3]string
	theArray[0] = "India"  // Assign a value to the first element
	theArray[1] = "Canada" // Assign a value to the second element
	theArray[2] = "Japan"  // Assign a value to the third element
	// for x := 0; x < len(theArray); x++ {
    //     fmt.Println(strings.Count(theArray[x],"a"))
	// 	append(theCounterArray, strings.Count(theArray[x],"a"))
    // }
	// fmt.Println(theCounterArray)
	fmt.Println("Before sorting  " ,theArray)
	sort.SliceStable(theArray, func(i, j int) bool {
		return strings.Count(theArray[i],"a") < strings.Count(theArray[j],"a")
	})
	fmt.Println(theArray)
}
func main() {
    fmt.Println("!... Hello World ...!")
	q_1()
}