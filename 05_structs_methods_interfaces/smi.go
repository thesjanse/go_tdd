package smi

type Rectangle struct {
	width  float64
	height float64
}

func Perimeter(r Rectangle) float64 {
	return (r.width + r.height) * 2
}

func Area(r Rectangle) float64 {
	return r.width * r.height
}
