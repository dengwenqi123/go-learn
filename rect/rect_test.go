package rect

import (
	"testing"
	"fmt"
)

func TestArea_Area(t *testing.T) {
	ar := Area{100,200}
	val := ar.Area()
	fmt.Println(val)

}

func TestPrint(t *testing.T) {
	ar := Area{100,400}
	val := ar.Area()
	Print(val)
}
