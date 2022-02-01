package main

import (
	"fmt"
)

const (
	pi = 3.1415
	first = iota
	second = 2 << iota
	third
) 


func main() {
	fmt.Println(pi)

	fmt.Println(first, second, third)

	
	const c = 3
	fmt.Println(c + 3, c + 1.2)
	
	fmt.Println("Hello from module gophers")

	var firstName *string = new(string)
	*firstName = "Arthur"
	fmt.Println(*firstName)

	lastName := "Lancelot"
	fmt.Println(lastName)

	ptr := &lastName
	fmt.Println(ptr, *ptr)

	lastName = "Tricia"
	fmt.Println(ptr, *ptr)

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arrfoo := [3]int {1, 2, 3}
	fmt.Println(arrfoo)

	slice := arr[:]
	fmt.Println(slice)

	arr[1]  = 42
	slice[2]= 27
	fmt.Println(arr, slice)

	slicefoo := []int {1, 2, 3}
	fmt.Println(slicefoo)

	slicefoo = append(slice, 4, 42, 48)
	fmt.Println(slicefoo)

	s2 := slicefoo[1:]
	
	s3 := slicefoo[:2]
	
	s4 := slicefoo[2:3]
	fmt.Println(s2, s3, s4)

	m := map[string]int{"foo":42}
	fmt.Println(m, m["foo"])

	m["foo"] = 23
	fmt.Println(m["foo"])

	delete(m, "foo")
	fmt.Println(m)
}