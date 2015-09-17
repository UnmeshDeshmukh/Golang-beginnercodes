package main

import "testing"
import "math"


func Test_Perimeter(t *testing.T){
  var num float64
  Rect := Rectangle{length:6, breadth:6}	    
  num = Rect.perimeter()
  if num !=24 {
    t.Error("Expected 24 got ", num)
  }
 }


func Test_PerimeterRect(t *testing.T){
  var num float64
  Rect := Rectangle{length:15, breadth:12}	    
  num = Rect.perimeter()
  if num !=54 {
    t.Error("Expected 54 got ", num)
  }
 
}



func Test_PerimeterRect1(t *testing.T){
  var num float64
  Rect := Rectangle{length:10, breadth:6}	    
  num = Rect.perimeter()
  if num !=32 {
    t.Error("Expected 32 got ", num)
  } 
}

func Test_PerimeterCircle(t *testing.T){
  var num float64
  var res float64
  varcircle := Circle{radius: 7}
  num = varcircle.perimeter()
  res = 2 * math.Pi * 7
  if num !=res {
    t.Error("Expected 22 got ", num)
 }
}


func Test_PerimeterCircle1(t *testing.T){
  var num float64
  var res float64
  varcircle := Circle{radius: 20}
  num = varcircle.perimeter()
  res = 2 * math.Pi * 20
  if num !=res {
    t.Error("Expected 22 got ", num)
 }
}

func Test_PerimeterCircle2(t *testing.T){
  var num float64
  var res float64
  varcircle := Circle{radius: 10}
  num = varcircle.perimeter()
  res = 2 * math.Pi * 10
  if num !=res {
    t.Error("Expected 22 got ", num)
 }
}
