package solver

type Solver interface {
	GetDay() int
	GetTestInput() string
	Parse(string)
	Part1() (int, error)
	Part2() (int, error)
}
