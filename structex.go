package main

import "fmt"

type Cat struct{
  name string
  color string
  age int 
}

func main(){
  var kit Cat

  kit.name="Bravo"
  kit.color="black & white"
  kit.age=4

 details(kit) 
}

func details(kitten Cat){
  fmt.Println("Name of Cat is :",kitten.name)
  fmt.Println("Color of the cat is :",kitten.color)
  fmt.Println("Age of Cat is :",kitten.age)
  
}

