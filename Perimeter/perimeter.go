package main
//import "fmt"
import "math"

type Shape interface {
    perimeter() float64
}

type Circle struct {
    radius float64
}

type Rectangle struct{
    length float64
    breadth float64
}
func (sqside Rectangle) perimeter() float64{
    return 2 * (sqside.length + sqside.breadth)
}
func (varcircle Circle) perimeter() float64{
    return 2 * math.Pi * varcircle.radius
}


/*func print (shape Shape){
    fmt.Println(shape.perimeter())
}
func main() {
   varcircle := Circle{radius: 7}
   print(varcircle)
   side := Rectangle{length: 5, breadth:6}
   print(side)
}*/
