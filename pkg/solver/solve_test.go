package solver

import (
	"flag"
	"fmt"
	"github.com/stretchr/testify/require"
	"go-2x2-solver/pkg/cube"
	"testing"
)

var batches = flag.Int("batches", 4, "number of batches of tests")
var batchSize = flag.Int("batch_size", 100, "number of tests in each batch")
var scrambleLength = flag.Int("scramble_length", 20, "length of scramble")

func Test_Solve(t *testing.T) {
	for i := 0; i < *batches; i++ {
		t.Run(fmt.Sprintf("batch %d", i+1), func(t *testing.T) {
			t.Parallel()

			for i := 0; i < *batchSize; i++ {
				scrambled := cube.GetScrambled(*scrambleLength)

				solution, err := Solve(scrambled)

				require.NoError(t, err)
				solved := applyAlgorithm(scrambled, solution)
				require.True(t, solved.IsSolved())
			}
		})
	}
}

func Benchmark_Solve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scrambled := cube.GetScrambled(*scrambleLength)
		_, err := Solve(scrambled)
		require.NoError(b, err)
	}
}

func applyAlgorithm(c cube.Cube, algorithm Algorithm) cube.Cube {
	for _, move := range algorithm {
		c = cube.MakeMove(c, move)
	}
	return c
}
