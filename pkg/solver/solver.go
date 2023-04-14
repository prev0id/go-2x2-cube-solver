package solver

import (
	"errors"
	"go-2x2-solver/pkg/cube"
	"log"
)

const (
	forwardMaxDepth  = 6
	backwardMaxDepth = 5
)

type queueValue struct {
	cube      cube.Cube
	algorithm []string
}

type Solver struct {
	forwardMemory   map[cube.Cube][]string
	backwardMemory  map[cube.Cube][]string
	resultAlgorithm []string
}

func (s *Solver) Solve(c cube.Cube) ([]string, error) {
	if c.IsSolved() {
		return nil, nil
	}

	// forward bfs
	s.forwardMemory = make(map[cube.Cube][]string)
	s.forwardMemory[c] = nil
	queue := make([]queueValue, 0, 1)
	queue = append(queue, queueValue{cube: c, algorithm: nil})
	resultForward := s.forwardBFS(queue, 1)

	// backward bfs
	if !resultForward {
		s.backwardMemory = make(map[cube.Cube][]string)
		queue := make([]queueValue, 0, 24)
		for _, solvedCube := range cube.SolvedCubes {
			s.backwardMemory[solvedCube] = nil
			queue = append(queue, queueValue{cube: c, algorithm: nil})
		}
		if err := s.backwardBFS(queue, 1); err != nil {
			return nil, err
		}
	}
	return s.resultAlgorithm, nil
}

func (s *Solver) forwardBFS(bfsQueue []queueValue, depth int) bool {
	if depth > forwardMaxDepth {
		return false
	}
	log.Printf("forwardBFS: depth [%d]\n", depth)

	forwardMoves := map[string]cube.Move{
		"R":  cube.MoveR,
		"R2": cube.MoveR2,
		"R'": cube.MoveRPrime,
		"U":  cube.MoveU,
		"U2": cube.MoveU2,
		"U'": cube.MoveUPrime,
		"F":  cube.MoveF,
		"F2": cube.MoveF2,
		"F'": cube.MoveFPrime,
	}

	// Test: predict capacity: len(bfsQueue)*9
	newQueue := make([]queueValue, 0, len(bfsQueue)*9)
	for _, value := range bfsQueue {
		prevCube, prevAlgorithm := value.cube, value.algorithm
		for moveName, move := range forwardMoves {
			newCube := cube.MakeMove(prevCube, move)
			if _, isExist := s.forwardMemory[newCube]; !isExist {
				if newCube.IsSolved() {
					s.resultAlgorithm = makeNewAlgorithm(prevAlgorithm, moveName)
					return true
				}
				// deep copy of algorithm slice
				newAlgorithm := makeNewAlgorithm(prevAlgorithm, moveName)
				// updating memory and queue
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
	log.Printf("backwardBFS: depth [%d]\n", depth)

	backwardMoves := map[string]cube.Move{
		"R":  cube.MoveRPrime,
		"R2": cube.MoveR2,
		"R'": cube.MoveR,
		"U":  cube.MoveUPrime,
		"U2": cube.MoveU2,
		"U'": cube.MoveU,
		"F":  cube.MoveFPrime,
		"F2": cube.MoveF2,
		"F'": cube.MoveF,
	}

	newQueue := make([]queueValue, 0, len(bfsQueue)*9)
	for _, value := range bfsQueue {
		prevCube, prevAlgorithm := value.cube, value.algorithm
		for moveName, move := range backwardMoves {
			newCube := cube.MakeMove(prevCube, move)
			if _, isExist := s.forwardMemory[newCube]; isExist {
				// extending forward algorithm with reversed backward algorithm
				s.resultAlgorithm = s.forwardMemory[newCube]
				backwardAlgorithm := makeNewAlgorithm(prevAlgorithm, moveName)
				reverseSlice(backwardAlgorithm)
				s.resultAlgorithm = append(s.resultAlgorithm, backwardAlgorithm...)
				return nil
			}
			if _, isExist := s.forwardMemory[newCube]; !isExist {
				// deep copy of algorithm slice
				newAlgorithm := makeNewAlgorithm(prevAlgorithm, moveName)
				// updating memory and queue
				s.backwardMemory[newCube] = newAlgorithm
				newQueue = append(newQueue, queueValue{newCube, newAlgorithm})
			}
		}
	}
	return s.backwardBFS(newQueue, depth+1)
}

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
