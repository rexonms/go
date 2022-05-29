package greet

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello from github.com/rexonms/go/greet"
	if got != want {
		t.Errorf("Hello(): got %v, want %v", got, want)
	}
}