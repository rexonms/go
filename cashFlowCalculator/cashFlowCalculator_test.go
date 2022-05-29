package cashFlowCalculator

import "testing"


func TestGetRentalCashFlow(t *testing.T) {
	gotA := GetRentalCashFlow(5000, 3000)
	wantA := 2000
	if gotA != wantA {
		t.Errorf("A GetRentalCashFlow(): got %v, want %v", gotA, wantA)
	}

	gotB := GetRentalCashFlow(3000, 5000)
	wantB := -2000
	if gotB != wantB {
		t.Errorf("B GetRentalCashFlow(): got %v, want %v", gotB, wantB)
	}
}