package main

import (
	"fmt"

	"github.com/rexonms/go/greet"
	"github.com/rexonms/go/property"
)


func main() {
	greet.Hello()
	fmt.Println(property.GetBathRoomCount(2, 1))
}