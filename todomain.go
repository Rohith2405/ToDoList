package main

import "fmt"

func main(){
  
  var elems int

  fmt.Println("Number of Activities: ")

  fmt.Scan(&elems)

	var slice = make([]string,elems) 
  for i := 0; i < elems; i++ {
     fmt.Println("Enter the activity",i+1)
  	fmt.Scan(&slice[i])
  }
  
  fmt.Println("Today Activity:")
	for i := 0; i < elems; i++ {
    fmt.Println("Activity: ",slice[i])
	}
	var j int
  for{
  fmt.Println("Completed Activity Number: ")
  fmt.Scan(&j)
  if(j<1 || j>elems){
      fmt.Println("Enter valid activity number")
      continue
  }
	slice = removeindex(slice, j-1)
	elems--
	if(elems == 0){
	    fmt.Println("No activities left. Congrats! you finished all Activities.")
	    break
	}
	fmt.Println("Activities Left: ")
		for i := 0; i < elems; i++ {
    fmt.Println("Activity: ",slice[i])
	}
  }
}

func removeindex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
