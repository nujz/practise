package main

import "sort"
import "fmt"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	res := make([][]int, len(people))
	for _, person := range people {
		score := person[1]
		for i := range res {
			if res[i] == nil {
				if score == 0 {
					res[i] = person
					break
				}
				score--
			}
		}
	}
	return res
}

func main() {
	a := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	b := reconstructQueue(a)
	fmt.Println(b)
}
