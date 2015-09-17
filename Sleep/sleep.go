package main
 import (
	"fmt"
	"time"
)

func Sleep(timer int){
	if timer<0{
		fmt.Println("Error !")
		return	 
        }
	<- time.After(time.Second * time.Duration(timer))
}

/*func main(){
	fmt.Println("Time Before:", time.Now().Format(time.RFC1123))
	Sleep(4);
	fmt.Println("Time After:", time.Now().Format(time.RFC1123))
}*/
