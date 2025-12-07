package internal

type SolverFunc func(string) string

var solvers = map[int]SolverFunc{}

func Register(day int, solver SolverFunc) {
	solvers[day] = solver
}

func GetSolver(day int) SolverFunc {
	return solvers[day]
}
