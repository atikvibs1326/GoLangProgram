package main

import "fmt"

func main(){
  control()
}

func control() {
	//Q: print 1 to 10 but skip 7
	for i := 0; i <= 10; i++ {
		if i == 7 {
			continue
		}
    fmt.Println(i)
	}
	
}
