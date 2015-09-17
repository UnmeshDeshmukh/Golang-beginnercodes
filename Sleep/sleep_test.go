package main

import "testing"
import "time"

func Test_Sleep(t *testing.T){
    timenow := time.Now()
    Sleep(3)
    timelater := time.Now()	
    if timelater.Sub(timenow) <=3 {
    t.Error("Some Error Occured")
  }
}
