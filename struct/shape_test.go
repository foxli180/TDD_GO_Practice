package _struct

import "testing"

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(shape Shape, want float64, t *testing.T) {
		got := shape.Perimeter();
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("Check rectangle perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 40.0
		checkPerimeter(rectangle, want, t)
	})

	t.Run("Check cycle perimeter", func(t *testing.T) {
		cycle := Cycle{5.0}
		want := 31.41592653589793
		checkPerimeter(cycle, want, t)
	})

}

func TestArea(t *testing.T) {

	checkArea := func(shape Shape, want float64, t *testing.T) {
		t.Helper()

		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		want := 72.0

		checkArea(rectangle, want, t)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Cycle{10}
		want := 314.1592653589793
		checkArea(circle, want, t)
	})
}

func TestArea2(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72},
		{name: "Cycle", shape: Cycle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Width: 12, Height: 6}, hasArea: 36},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})

	}
}
