package main

import "testing"

func Test_Fibonacci(t *testing.T){
  var num float64
  num = fib(-1)
  if num !=2 {
    t.Error("Expected 2 got ", num)
  }
}


func Test_Fibonacci1(t *testing.T){
  var num float64
  num = fib(7)
  if num !=13 {
    t.Error("Expected 13 got ", num)
  }
}


func Test_Fibonacci2(t *testing.T){
  var num float64
  num = fib(6)
  if num !=8 {
    t.Error("Expected 5 got ", num)
  }
}
