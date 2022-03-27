package main

import (
	"bitwise-operations-ex02/set"
	"fmt"
)

func main() {
	s := set.Set(1)
	s1 := set.Set(1)
	s2 := set.Set(2)
	fmt.Println(set.String(s))
	fmt.Println(set.IsEmpty(s))
	fmt.Println(set.Len(s))
	fmt.Println(set.Elements(s))
	fmt.Println(set.Add(s, 2))
	fmt.Println(set.Contains(s, 1))
	fmt.Println(set.Remove(s, 2))
	fmt.Println(set.Union(s1, s2))
	fmt.Println(set.Intersection(s1, s2))
	fmt.Println(set.Difference(s1, s2))
	fmt.Println(set.Subtract(s1, s2))
}
