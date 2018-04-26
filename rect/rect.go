package rect

import "fmt"

type Area struct {
	X,Y int
}
func (a *Area)Area() int {
	return a.X * a.Y
}

func Print(val int){
	fmt.Printf("this is the area %d",val)
}
