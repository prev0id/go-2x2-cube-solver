// Package solver implements optimal solver for 2x2x2 Rubik's cube
package solver

import (
	"go-2x2-solver/pkg/cube"
)

type solver struct {
	allMoves []cube.Move

	forwardMemory   map[cube.Cube]Algorithm
	forwardQueue    []queueValue
	forwardSolution Algorithm

	backwardMemory   map[cube.Cube]Algorithm
	backwardQueue    []queueValue
	backwardSolution Algorithm
}

type Algorithm = []cube.Move

type queueValue struct {
	cube      cube.Cube
	algorithm Algorithm
}

const (
	forwardMaxDepth  = 6
	backwardMaxDepth = 5
)

func (s *solver) forwardBFS(depth int) bool {
	if depth > forwardMaxDepth {
		return false
	}

	// capacity will increase ~6 times
	newQueue := make([]queueValue, 0, len(s.forwardQueue)*6)

	for _, value := range s.forwardQueue {
		prevCube, prevAlgorithm := value.cube, value.algorithm
		for _, move := range s.allMoves {
			newCube := cube.MakeMove(prevCube, move)
			if _, isExist := s.forwardMemory[newCube]; !isExist {
				if newCube.IsSolved() {
					s.forwardSolution = makeNewAlgorithm(prevAlgorithm, move)
					return true
				}

				newAlgorithm := makeNewAlgorithm(prevAlgorithm, move)
				s.forwardMemory[newCube] = newAlgorithm
				newQueue = append(newQueue, queueValue{newCube, newAlgorithm})
			}
		}
	}
	s.forwardQueue = newQueue
	return s.forwardBFS(depth + 1)
}

func (s *solver) backwardBFS(depth int) bool {
	if depth > backwardMaxDepth {
		return false
	}
	// capacity will increase ~6 times
	newQueue := make([]queueValue, 0, len(s.backwardQueue)*6)

	for _, value := range s.backwardQueue {
		prevCube, prevAlgorithm := value.cube, value.algorithm
		for _, move := range s.allMoves {
			newCube := cube.MakeMove(prevCube, move)

			if _, isExist := s.forwardMemory[newCube]; isExist {
				s.forwardSolution = s.forwardMemory[newCube]
				s.backwardSolution = makeNewAlgorithm(prevAlgorithm, move)
				return true
			}

			if _, isExist := s.backwardMemory[newCube]; !isExist {
				newAlgorithm := makeNewAlgorithm(prevAlgorithm, move)
				s.backwardMemory[newCube] = newAlgorithm
				newQueue = append(newQueue, queueValue{newCube, newAlgorithm})
			}
		}
	}
	s.backwardQueue = newQueue
	return s.backwardBFS(depth + 1)
}

// calcAlgorithm constructing from solutions of forward and backwards BFS's final algorithm
func (s *solver) calcAlgorithm() Algorithm {
	result := make(Algorithm, 0, len(s.backwardSolution)+len(s.forwardSolution))
	result = append(result, s.forwardSolution...)

	reverseThisMove := map[cube.Move]cube.Move{
		cube.R:      cube.RPrime,
		cube.R2:     cube.R2,
		cube.RPrime: cube.R,
		cube.U:      cube.UPrime,
		cube.U2:     cube.U2,
		cube.UPrime: cube.U,
		cube.F:      cube.FPrime,
		cube.F2:     cube.F2,
		cube.FPrime: cube.F,
	}
	for idx := len(s.backwardSolution) - 1; idx >= 0; idx-- {
		result = append(result, reverseThisMove[s.backwardSolution[idx]])
	}
	return result
}

// makeNewAlgorithm appends move to deep copy of algorithm
func makeNewAlgorithm(algorithm Algorithm, move cube.Move) Algorithm {
	newAlgorithm := make([]cube.Move, 0, len(algorithm)+1)
	newAlgorithm = append(newAlgorithm, algorithm...)
	newAlgorithm = append(newAlgorithm, move)
	return newAlgorithm
}
