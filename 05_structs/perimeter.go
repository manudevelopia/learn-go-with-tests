package structs

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Width)
}

func Area(rectangle Rectangle) interface{} {
	return rectangle.Width * rectangle.Height
}
