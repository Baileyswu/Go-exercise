package main

import "fmt"

func main() {
	fmt.Println("func c = ", c())
	fmt.Println(func (a int) int{ return a*a} (5))
}

func c() (i int) {
	fmt.Println("start c")
	defer func() { 
		fmt.Println("in c")
		i++ 
	}()
	fmt.Println("end c")
    return 1
}