package main

import (
	"fmt"

	"github.com/rexonms/go/cashFlowCalculator"
	"github.com/rexonms/go/property"
)


func main() {

	// Property
	fmt.Println("*** Property ***")
	fmt.Println("GetBathRoomCount:", property.GetBathRoomCount(2, 1))

	// CashFlowCalculator
	fmt.Println("*** CashFlowCalculator ***")
	fmt.Println("GetRentalCashFlow:",cashFlowCalculator.GetRentalCashFlow(5000, 3000))

}