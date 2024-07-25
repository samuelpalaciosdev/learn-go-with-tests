package structs_interfaces

import "testing"

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{2.75, 5}
		got := rectangle.Area()
		want := 13.75
		if got != want {
			t.Errorf("got %2.f want %2.f", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

// Perimeter of a rectangle
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %2.f want %2.f", got, want)
	}
}
