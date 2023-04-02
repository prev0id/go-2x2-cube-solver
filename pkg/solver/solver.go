package solver

import (
	"errors"
	"log"
)

const (
	forwardMaxDepth  = 6
	backwardMaxDepth = 5
)

var (
	ErrInvalidCube = errors.New("unsolvable cube")

	forwardMoves = map[string][24]int{
		MoveR:      RMoveIndexes,
		MoveR2:     R2MoveIndexes,
		MoveRPrime: RPrimeMoveIndexes,
		MoveU:      UMoveIndexes,
		MoveU2:     U2MoveIndexes,
		MoveUPrime: UPrimeMoveIndexes,
		MoveF:      FMoveIndexes,
		MoveF2:     F2MoveIndexes,
		MoveFPrime: FPrimeMoveIndexes}

	backwardMoves = map[string][24]int{
		MoveR:      RPrimeMoveIndexes,
		MoveR2:     R2MoveIndexes,
		MoveRPrime: RMoveIndexes,
		MoveU:      UPrimeMoveIndexes,
		MoveU2:     U2MoveIndexes,
		MoveUPrime: UMoveIndexes,
		MoveF:      FPrimeMoveIndexes,
		MoveF2:     F2MoveIndexes,
		MoveFPrime: FMoveIndexes}
)

type queueValue struct {
	cube      Cube
	algorithm []string
}

type Solver struct {
	forwardMemory   map[Cube][]string
	backwardMemory  map[Cube][]string
	resultAlgorithm []string
}

func (s *Solver) Solve(cube Cube) ([]string, error) {
	if cube.IsSolved() {
		return nil, nil
	}

	// forward bfs
	s.forwardMemory = make(map[Cube][]string)
	s.forwardMemory[cube] = nil
	queue := make([]queueValue, 0, 1)
	queue = append(queue, queueValue{cube: cube, algorithm: nil})
	resultForward := s.forwardBFS(queue, 1)

	// backward bfs
	if !resultForward {
		s.backwardMemory = make(map[Cube][]string)
		queue := make([]queueValue, 0, 24)
		for _, solvedState := range SolvedStickers {
			cube := Cube{}
			cube.SetStickers(solvedState)
			s.backwardMemory[cube] = nil
			queue = append(queue, queueValue{cube: cube, algorithm: nil})
		}
		if err := s.backwardBFS(queue, 1); err != nil {
			return nil, err
		}
	}
	return s.resultAlgorithm, nil
}

func (s *Solver) getAllSolvedStates() {
	for _, solvedState := range SolvedStickers {
		cube := Cube{}
		cube.SetStickers(solvedState)
		s.backwardMemory[cube] = nil
	}
}

func (s *Solver) forwardBFS(bfsQueue []queueValue, depth int) bool {
	if depth > forwardMaxDepth {
		return false
	}
	log.Printf("forwardBFS: depth [%d]\n", depth)

	// Test: predict capacity: len(bfsQueue)*9
	newQueue := make([]queueValue, 0, len(bfsQueue)*9)
	for _, value := range bfsQueue {
		prevCube, prevAlgorithm := value.cube, value.algorithm
		for moveName, move := range forwardMoves {
			newCube := prevCube.MakeMove(move)
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
		return ErrInvalidCube
	}
	log.Printf("backwardBFS: depth [%d]\n", depth)

	newQueue := make([]queueValue, 0, len(bfsQueue)*9)
	for _, value := range bfsQueue {
		prevCube, prevAlgorithm := value.cube, value.algorithm
		for moveName, move := range backwardMoves {
			newCube := prevCube.MakeMove(move)
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
