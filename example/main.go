package main

import (
	"fmt"

	"github.com/rexonms/go/cashFlowCalculator"
	"github.com/rexonms/go/greet"
	"github.com/rexonms/go/property"
)


func main() {
	// Greet
	fmt.Println("*** Greet ***")
	fmt.Println("Hello:", greet.Hello())	
	
	// Property
	fmt.Println("*** Property ***")
	fmt.Println("GetBathRoomCount:", property.GetBathRoomCount(2, 1))

	// CashFlowCalculator
	fmt.Println("*** CashFlowCalculator ***")
	fmt.Println("GetRentalCashFlow:",cashFlowCalculator.GetRentalCashFlow(5000, 3000))

}