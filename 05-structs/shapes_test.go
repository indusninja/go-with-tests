package _5_structs

import "testing"

func TestPerimeter(t *testing.T) {

	verification := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("perimeter of a rectangle with sides - 10 and 10", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		verification(t, rectangle.Perimeter(), 40.0)
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 6, Height: 10}, hasArea: 60.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

}
