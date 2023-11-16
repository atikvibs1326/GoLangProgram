package main

import "fmt"

func main() {
	var myarr = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		fmt.Println(myarr[i]) 
	}
  myslice :=myarr[0:4]
  fmt.Println(myslice)
}
