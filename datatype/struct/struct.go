package main

type struc struct {
	field1 int
	field2 string
}

// Metod for structure.
func (s *struc) function(i int, str string) {
	s.field1 = i
	s.field2 = str
}

func main() {
	s := struc{0, ""}
	s.function(1, "some")
}
