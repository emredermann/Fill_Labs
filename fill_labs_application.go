package main

import(
	"fmt"
	"strings"
	"unicode/utf8"
	"sort"
	)
	/*
	Q1) Write a function that sorts a bunch of words by the number of character “a”s within the
		word (decreasing order). If some words contain the same amount of character “a”s then you
		need to sort those words by their lengths.
	*/
	
	

func q_1(in [] string) [] string{
	sort.Slice(in, func(i, j int) bool {
		s1, s2 := in[i], in[j]
		count1, count2 := strings.Count(s1, "a"), strings.Count(s2, "a")
		if count1 != count2 {
			return count1 > count2
		}
		return utf8.RuneCountInString(s1) > utf8.RuneCountInString(s2)
	})
	return in
}
/*
	Q2) Write a recursive function which takes one integer parameter. Please bear in mind that
		finding the algorithm needed to generate the output below is the main point of the question.
	*/
	func q_2(in int) {
		if in == 1 {
			return 
		}
		q_2(in/2)
		fmt.Println(in)
	}
/*
	Q3) Write a function which takes one parameter as an array/list. Find most repeated data
		within a given array.
	*/
	func q_3(arr [] string) string{
			wc := make(map[string]int)
			for _, word := range arr {
				_, matched := wc[word]
				if matched {
					wc[word] += 1
				} else {
					wc[word] = 1
				}
			}
			ref_key := 0
			ref_element :=""
			for element, key := range wc {
				if key > ref_key{
					ref_key = key
					ref_element  = element
				}
			}
			return ref_element
	}

func main() {
	// in := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	// arr := []string{"apple","pie","apple","red","red","red"}
	// fmt.Println(q_1(in))
	q_2(9)
	// fmt.Println(q_3(arr))

}