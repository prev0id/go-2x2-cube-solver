package cube

import "math/rand"

type Move int

const (
	R Move = iota
	R2
	RPrime
	U
	U2
	UPrime
	F
	F2
	FPrime
)

func MakeMove(cube Cube, move Move) Cube {
	newCube := Cube{}
	for newIdx, prevIdx := range moves[move] {
		newCube[newIdx] = cube[prevIdx]
	}
	return newCube
}

func GetScrambled(scrambleLength int) Cube {
	var possibleMoves = []Move{R, R2, RPrime, U, U2, UPrime, F, F2, FPrime}
	cube := getRandomFromSlice(GetSolvedCubes())

	for j := 0; j < scrambleLength; j++ {
		cube = MakeMove(cube, getRandomFromSlice(possibleMoves))
	}
	return cube
}

func getRandomFromSlice[Type any](slice []Type) Type {
	return slice[rand.Intn(len(slice))]
}

// moves matches the movement to the change of the sticker indexes in the cube
var moves = [][]int{
	R:      {0, 9, 2, 11, 4, 5, 6, 7, 8, 21, 10, 23, 14, 12, 15, 13, 3, 17, 1, 19, 20, 18, 22, 16},
	R2:     {0, 21, 2, 23, 4, 5, 6, 7, 8, 18, 10, 16, 15, 14, 13, 12, 11, 17, 9, 19, 20, 1, 22, 3},
	RPrime: {0, 18, 2, 16, 4, 5, 6, 7, 8, 1, 10, 3, 13, 15, 12, 14, 23, 17, 21, 19, 20, 9, 22, 11},
	U:      {2, 0, 3, 1, 8, 9, 6, 7, 12, 13, 10, 11, 16, 17, 14, 15, 4, 5, 18, 19, 20, 21, 22, 23},
	U2:     {3, 2, 1, 0, 12, 13, 6, 7, 16, 17, 10, 11, 4, 5, 14, 15, 8, 9, 18, 19, 20, 21, 22, 23},
	UPrime: {1, 3, 0, 2, 16, 17, 6, 7, 4, 5, 10, 11, 8, 9, 14, 15, 12, 13, 18, 19, 20, 21, 22, 23},
	F:      {0, 1, 7, 5, 4, 20, 6, 21, 10, 8, 11, 9, 2, 13, 3, 15, 16, 17, 18, 19, 14, 12, 22, 23},
	F2:     {0, 1, 21, 20, 4, 14, 6, 12, 11, 10, 9, 8, 7, 13, 5, 15, 16, 17, 18, 19, 3, 2, 22, 23},
	FPrime: {0, 1, 12, 14, 4, 3, 6, 2, 9, 11, 8, 10, 21, 13, 20, 15, 16, 17, 18, 19, 5, 7, 22, 23},
}
