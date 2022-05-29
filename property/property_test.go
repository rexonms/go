package property

import "testing"

func TestGetBathRoomCount(t *testing.T) {
	gotA := GetBathRoomCount(1,0)
	wantA := float32(1)
	if gotA != wantA {
		t.Errorf("A GetBathRoomCount(): got %v, want %v", gotA, wantA)
	}

	gotB := GetBathRoomCount(2,1)
	wantB := float32(2.5)
	if gotB != wantB {
		t.Errorf("B GetBathRoomCount(): got %v, want %v", gotB, wantB)
	}

	gotC := GetBathRoomCount(2,0)
	wantC := float32(2)
	if gotC != wantC {
		t.Errorf("C GetBathRoomCount(): got %v, want %v", gotC, wantC)
	}

	gotD := GetBathRoomCount(0,0)
	wantD := float32(0)
	if gotD != wantD {
		t.Errorf("D GetBathRoomCount(): got %v, want %v", gotD, wantD)
	}
}