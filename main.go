package main

import (
    "fmt"
	"strings"
	"sort"
)
func main() {
    greeting := "Hello my friends!"
	fmt.Println(strings.Contains(greeting,"friends"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "my"))
	fmt.Println(strings.Split(greeting, "friends"))
	ages := []int{30, 25, 22}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 30)
	fmt.Println(index)
	names := []string {"Caruzo", "Almofadinha", "Gobas"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "Caruzo"))
}
		
		