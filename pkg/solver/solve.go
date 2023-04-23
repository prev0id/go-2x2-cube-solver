package solver

import (
	"errors"
	"go-2x2-solver/pkg/cube"
)

func Solve(c cube.Cube) (Algorithm, error) {
	if c.IsSolved() {
		return nil, nil
	}

	s := solver{
		allMoves: []cube.Move{
			cube.R, cube.R2, cube.RPrime,
			cube.U, cube.U2, cube.UPrime,
			cube.F, cube.F2, cube.FPrime},

		forwardMemory: make(map[cube.Cube][]cube.Move),
		forwardQueue:  make([]queueValue, 0, 1),

		backwardMemory: make(map[cube.Cube][]cube.Move),
		backwardQueue:  make([]queueValue, 0, 24),
	}

	// forward bfs
	s.forwardMemory[c] = nil
	s.forwardQueue = append(s.forwardQueue, queueValue{cube: c, algorithm: nil})
	successForward := s.forwardBFS(1)

	// backward bfs
	if !successForward {
		for _, solvedCube := range cube.GetSolvedCubes() {
			s.backwardMemory[solvedCube] = nil
			s.backwardQueue = append(s.backwardQueue, queueValue{cube: solvedCube, algorithm: nil})
		}
		if !s.backwardBFS(1) {
			return nil, errors.New("unsolvable cube")
		}
	}

	return s.calcAlgorithm(), nil
}
