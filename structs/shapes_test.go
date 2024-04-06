package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("perimeter of a rectangle with sides - 10 and 10", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		assertCorrectMessage(t, rectangle.Perimeter(), 40.0)
	})
}

func TestArea(t *testing.T) {
	t.Run("area of a rectangle with sides - 6 and 10", func(t *testing.T) {
		rectangle := Rectangle{6.0, 10.0}
		assertCorrectMessage(t, rectangle.Area(), 60.0)
	})
	t.Run("area of a circle with radius - 10", func(t *testing.T) {
		circle := Circle{10.0}
		assertCorrectMessage(t, circle.Area(), 314.1592653589793)
	})
}

func assertCorrectMessage(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
