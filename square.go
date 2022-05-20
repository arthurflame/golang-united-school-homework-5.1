package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

//func (receiver) End() Point {
//	// implement me
//}

func (s *Square) Area() uint {
	return s.a * s.a
}

func (s *Square) Perimeter() uint {
	return s.a * 4
}
