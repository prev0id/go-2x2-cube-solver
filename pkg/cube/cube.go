// Package cube implements 2x2x2 Rubik's cube (pocket cube)
package cube

type Cube [24]Sticker
type Sticker int8

// numbering of cube stickers:
//
//		 +--+--+
//		 | 0| 1|
//		 +--+--+
//		 | 2| 3|
// +--+--+--+--+--+--+--+--+
// | 4| 5| 8| 9|12|13|16|17|
// +--+--+--+--+--+--+--+--+
// | 6| 7|10|11|14|15|18|19|
// +--+--+--+--+--+--+--+--+
//		 |20|21|
//		 +--+--+
//		 |22|23|
//		 +--+--+

// IsSolved returns true if the cube can be considered solved
func (c *Cube) IsSolved() bool {
	_, isExist := solvedCubes[*c]
	return isExist
}

// GetSolvedCubes returns slice of all rotations of the solved cube
func GetSolvedCubes() []Cube {
	cubes := make([]Cube, 0, len(solvedCubes))
	for cube := range solvedCubes {
		cubes = append(cubes, cube)
	}
	return cubes
}

// lookup table for all rotations of the solved cube
var solvedCubes = map[Cube]struct{}{
	{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5}: {},
	{3, 3, 3, 3, 0, 0, 0, 0, 2, 2, 2, 2, 5, 5, 5, 5, 4, 4, 4, 4, 1, 1, 1, 1}: {},
	{5, 5, 5, 5, 3, 3, 3, 3, 2, 2, 2, 2, 1, 1, 1, 1, 4, 4, 4, 4, 0, 0, 0, 0}: {},
	{1, 1, 1, 1, 5, 5, 5, 5, 2, 2, 2, 2, 0, 0, 0, 0, 4, 4, 4, 4, 3, 3, 3, 3}: {},
	{5, 5, 5, 5, 1, 1, 1, 1, 4, 4, 4, 4, 3, 3, 3, 3, 2, 2, 2, 2, 0, 0, 0, 0}: {},
	{3, 3, 3, 3, 5, 5, 5, 5, 4, 4, 4, 4, 0, 0, 0, 0, 2, 2, 2, 2, 1, 1, 1, 1}: {},
	{0, 0, 0, 0, 3, 3, 3, 3, 4, 4, 4, 4, 1, 1, 1, 1, 2, 2, 2, 2, 5, 5, 5, 5}: {},
	{1, 1, 1, 1, 0, 0, 0, 0, 4, 4, 4, 4, 5, 5, 5, 5, 2, 2, 2, 2, 3, 3, 3, 3}: {},
	{0, 0, 0, 0, 4, 4, 4, 4, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 5, 5, 5, 5}: {},
	{2, 2, 2, 2, 0, 0, 0, 0, 1, 1, 1, 1, 5, 5, 5, 5, 3, 3, 3, 3, 4, 4, 4, 4}: {},
	{5, 5, 5, 5, 2, 2, 2, 2, 1, 1, 1, 1, 4, 4, 4, 4, 3, 3, 3, 3, 0, 0, 0, 0}: {},
	{4, 4, 4, 4, 5, 5, 5, 5, 1, 1, 1, 1, 0, 0, 0, 0, 3, 3, 3, 3, 2, 2, 2, 2}: {},
	{0, 0, 0, 0, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 1, 1, 1, 1, 5, 5, 5, 5}: {},
	{4, 4, 4, 4, 0, 0, 0, 0, 3, 3, 3, 3, 5, 5, 5, 5, 1, 1, 1, 1, 2, 2, 2, 2}: {},
	{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3, 2, 2, 2, 2, 1, 1, 1, 1, 0, 0, 0, 0}: {},
	{2, 2, 2, 2, 5, 5, 5, 5, 3, 3, 3, 3, 0, 0, 0, 0, 1, 1, 1, 1, 4, 4, 4, 4}: {},
	{2, 2, 2, 2, 1, 1, 1, 1, 5, 5, 5, 5, 3, 3, 3, 3, 0, 0, 0, 0, 4, 4, 4, 4}: {},
	{1, 1, 1, 1, 4, 4, 4, 4, 5, 5, 5, 5, 2, 2, 2, 2, 0, 0, 0, 0, 3, 3, 3, 3}: {},
	{4, 4, 4, 4, 3, 3, 3, 3, 5, 5, 5, 5, 1, 1, 1, 1, 0, 0, 0, 0, 2, 2, 2, 2}: {},
	{3, 3, 3, 3, 2, 2, 2, 2, 5, 5, 5, 5, 4, 4, 4, 4, 0, 0, 0, 0, 1, 1, 1, 1}: {},
	{2, 2, 2, 2, 3, 3, 3, 3, 0, 0, 0, 0, 1, 1, 1, 1, 5, 5, 5, 5, 4, 4, 4, 4}: {},
	{3, 3, 3, 3, 4, 4, 4, 4, 0, 0, 0, 0, 2, 2, 2, 2, 5, 5, 5, 5, 1, 1, 1, 1}: {},
	{4, 4, 4, 4, 1, 1, 1, 1, 0, 0, 0, 0, 3, 3, 3, 3, 5, 5, 5, 5, 2, 2, 2, 2}: {},
	{1, 1, 1, 1, 2, 2, 2, 2, 0, 0, 0, 0, 4, 4, 4, 4, 5, 5, 5, 5, 3, 3, 3, 3}: {},
}
