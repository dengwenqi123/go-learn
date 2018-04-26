package main

import (
	"github.com/dengwenqi123/go-learn/rect"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	a := rect.Area{20,30}
	b := a.Area()
	rect.Print(b)
	fmt.Println("main project")

	db,err := bolt.Open("my.db",0600,nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
