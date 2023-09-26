package shapes_test

import (
	"testing"

	si "github.com/Longoria-Yu/go-with-tests/struct_interface"
)

func TestPerimeter(t *testing.T) {
	rectangle := si.Rectangle{10.0, 10.0}
	got := si.Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := si.Rectangle{12.0, 6.0}
	got := si.Area(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
