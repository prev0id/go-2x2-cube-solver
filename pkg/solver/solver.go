// Package solver implements optimal solver for 2x2x2 Rubik's cube
package solver

import (
	"errors"
	"go-2x2-solver/pkg/cube"
)

func NewSolver() *Solver {
	return &Solver{}
}

type Solver struct {
	forwardMemory   map[cube.Cube][]string
	backwardMemory  map[cube.Cube][]string
	resultAlgorithm []string
}

const (
	forwardMaxDepth  = 6
	backwardMaxDepth = 5
)

type queueValue struct {
	cube      cube.Cube
	algorithm []string
}

func (s *Solver) Solve(c cube.Cube) ([]string, error) {
	if c.IsSolved() {
		return nil, nil
	}

	// forward bfs
	s.forwardMemory = make(map[cube.Cube][]string)
	s.forwardMemory[c] = nil
	forwardQueue := make([]queueValue, 0, 1)
	forwardQueue = append(forwardQueue, queueValue{cube: c, algorithm: nil})
	resultForward := s.forwardBFS(forwardQueue, 1)

	// backward bfs
	if !resultForward {
		s.backwardMemory = make(map[cube.Cube][]string)
		backwardQueue := make([]queueValue, 0, 24)
		for _, solvedCube := range cube.GetSolvedCubes() {
			s.backwardMemory[solvedCube] = nil
			backwardQueue = append(backwardQueue, queueValue{cube: solvedCube, algorithm: nil})
		}
		if err := s.backwardBFS(backwardQueue, 1); err != nil {
			return nil, err
		}
	}
	return s.resultAlgorithm, nil
}

func (s *Solver) forwardBFS(bfsQueue []queueValue, depth int) bool {
	if depth > forwardMaxDepth {
		return false
	}

	forwardMoves := map[string]cube.Move{
		"R":  cube.R,
		"R2": cube.R2,
		"R'": cube.RPrime,
		"U":  cube.U,
		"U2": cube.U2,
		"U'": cube.UPrime,
		"F":  cube.F,
		"F2": cube.F2,
		"F'": cube.FPrime,
	}

	// capacity will increase ~6 times
	newQueue := make([]queueValue, 0, len(bfsQueue)*6)

	for _, value := range bfsQueue {
		prevCube, prevAlgorithm := value.cube, value.algorithm
		for moveName, move := range forwardMoves {
			newCube := cube.MakeMove(prevCube, move)
			if _, isExist := s.forwardMemory[newCube]; !isExist {
				if newCube.IsSolved() {
					s.resultAlgorithm = makeNewAlgorithm(prevAlgorithm, moveName)
					return true
				}

				newAlgorithm := makeNewAlgorithm(prevAlgorithm, moveName)
				s.forwardMemory[newCube] = newAlgorithm
				newQueue = append(newQueue, queueValue{newCube, newAlgorithm})
			}
		}
	}
	return s.forwardBFS(newQueue, depth+1)
}

func (s *Solver) backwardBFS(bfsQueue []queueValue, depth int) error {
	if depth > backwardMaxDepth {
		return errors.New("unsolvable cube")
	}

	backwardMoves := map[string]cube.Move{
		"R":  cube.RPrime,
		"R2": cube.R2,
		"R'": cube.R,
		"U":  cube.UPrime,
		"U2": cube.U2,
		"U'": cube.U,
		"F":  cube.FPrime,
		"F2": cube.F2,
		"F'": cube.F,
	}

	// capacity will increase ~6 times
	newQueue := make([]queueValue, 0, len(bfsQueue)*6)

	for _, value := range bfsQueue {
		prevCube, prevAlgorithm := value.cube, value.algorithm
		for moveName, move := range backwardMoves {
			newCube := cube.MakeMove(prevCube, move)

			if _, isExist := s.forwardMemory[newCube]; isExist {
				// constructing a solution using an extension of the forward algorithm
				// with a reverse backward algorithm
				s.resultAlgorithm = s.forwardMemory[newCube]
				backwardAlgorithm := makeNewAlgorithm(prevAlgorithm, moveName)
				reverseSlice(backwardAlgorithm)
				s.resultAlgorithm = append(s.resultAlgorithm, backwardAlgorithm...)
				return nil
			}

			if _, isExist := s.backwardMemory[newCube]; !isExist {
				newAlgorithm := makeNewAlgorithm(prevAlgorithm, moveName)
				s.backwardMemory[newCube] = newAlgorithm
				newQueue = append(newQueue, queueValue{newCube, newAlgorithm})
			}
		}
	}
	return s.backwardBFS(newQueue, depth+1)
}

// makeNewAlgorithm appends moveName to deep copy of algorithm
func makeNewAlgorithm(algorithm []string, moveName string) []string {
	newAlgorithm := make([]string, 0, len(algorithm)+1)
	newAlgorithm = append(newAlgorithm, algorithm...)
	newAlgorithm = append(newAlgorithm, moveName)
	return newAlgorithm
}

func reverseSlice(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
