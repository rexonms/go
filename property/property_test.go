package property

import "testing"

func TestGetBathRoomCount(t *testing.T) {
	responseA := GetBathRoomCount(1,0);
	if responseA != 1 {
		t.Errorf("Expected 1 but got %v", responseA)
	}
	responseB := GetBathRoomCount(2,1);
	if responseB != 2.5 {
		t.Errorf("Expected 2.5 but got %v", responseB)
	}
	responseC := GetBathRoomCount(2,0);
	if responseC != 2 {
		t.Errorf("Expected 2 but got %v", responseC)
	}
	responseD := GetBathRoomCount(0,0);
	if responseD != 0 {
		t.Errorf("Expected 0 but got %v", responseD)
	}
}