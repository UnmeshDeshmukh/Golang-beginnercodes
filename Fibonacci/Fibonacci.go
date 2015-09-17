package main
import "fmt"

func fib( x  float64 ) float64 {

if x<0{
  fmt.Println("Error Entered Negative number")
  return x	
}
if x==0 {
    return 0
}
if x==1 {
    return 1
}

return fib(x - 1) + fib(x - 2)
}

