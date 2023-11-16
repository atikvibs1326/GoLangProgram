package main

import "fmt"

func main(){
  recursive(1)
}
func recursive (n int ) int {
  //base condition
  if(n==5){
    fmt.Println(5)
    return 5
  }
  
  //basic printing it
  fmt.Println(n)
  //calling the function itself
  return recursive(n+1)
}
