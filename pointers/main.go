package main

import "fmt"

func main() {
     //var foo int =23
     //var x int = 5
     var bar int = 42
     //var baz int = 69

     var ptr *int = &bar
     fmt.Println(*ptr)
}