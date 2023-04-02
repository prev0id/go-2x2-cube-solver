package solver

import (
	"errors"
	"math/rand"
)

type Cube struct {
	stickers [24]Sticker
}

func (c *Cube) New() {
	c.stickers = SolvedStickers[0]
}

func (c *Cube) SetStickers(newStickers [24]Sticker) {
	c.stickers = newStickers
}

func (c *Cube) GetStickers() [24]Sticker {
	return c.stickers
}

func (c *Cube) MakeMove(moveIndexes [24]int) Cube {
	newCube := Cube{}
	for newIdx, prevIdx := range moveIndexes {
		newCube.stickers[newIdx] = c.stickers[prevIdx]
	}
	return newCube
}

func (c *Cube) IsSolved() bool {
	for _, solvedState := range SolvedStickers {
		if c.stickers == solvedState {
			return true
		}
	}
	return false
}

func (c *Cube) ApplyAlgorithm(algorithm []string) (Cube, error) {
	resultCube := *c
	for _, moveName := range algorithm {
		moveIndexes, isExist := cubeMoves[moveName]
		if !isExist {
			return *c, errors.New("invalid algorithm token [" + moveName + "]")
		}
		resultCube = resultCube.MakeMove(moveIndexes)
	}
	return resultCube, nil
}

func getRandomFromSlice[Type any](slice []Type) Type {
	return slice[rand.Intn(len(slice))]
}

func (c *Cube) ScrambleThis(scrambleLength int) []string {
	var possibleMoves = []string{MoveR, MoveR2, MoveRPrime, MoveU, MoveU2, MoveUPrime, MoveF, MoveF2, MoveFPrime}
	c.SetStickers(getRandomFromSlice(SolvedStickers))
	scrambleAlgorithm := make([]string, 0, scrambleLength)
	for j := 0; j < scrambleLength; j++ {
		scrambleAlgorithm = append(scrambleAlgorithm, getRandomFromSlice(possibleMoves))
	}
	*c, _ = c.ApplyAlgorithm(scrambleAlgorithm)
	return scrambleAlgorithm
}

//for _, moveName := range algorithm {
//	switch moveName {
//	case MoveR:
//		resultCube = resultCube.MakeMove(RMoveIndexes)
//	case MoveR2:
//		resultCube = resultCube.MakeMove(R2MoveIndexes)
//	case MoveRPrime:
//		resultCube = resultCube.MakeMove(RPrimeMoveIndexes)
//	case MoveU:
//		resultCube = resultCube.MakeMove(UMoveIndexes)
//	case MoveU2:
//		resultCube = resultCube.MakeMove(U2MoveIndexes)
//	case MoveUPrime:
//		resultCube = resultCube.MakeMove(UPrimeMoveIndexes)
//	case MoveF:
//		resultCube = resultCube.MakeMove(FMoveIndexes)
//	case MoveF2:
//		resultCube = resultCube.MakeMove(F2MoveIndexes)
//	case MoveFPrime:
//		resultCube = resultCube.MakeMove(FPrimeMoveIndexes)
//	default:
//		return *c, errors.New("invalid algorithm token [" + moveName + "]")
//	}
//}
//return resultCube, nil
