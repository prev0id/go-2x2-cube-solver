package solver

import (
	"fmt"
	"go-2x2-solver/pkg/cube"
	"testing"
)

func TestSolver_Solve(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("batch #%d", i+1),
			func(t *testing.T) {
				t.Parallel()
				solver := NewSolver()
				for i := 0; i < 100; i++ {
					scrambled := cube.GetScrambled(20)
					_, err := solver.Solve(scrambled)
					if err != nil {
						t.Fatal(err.Error())
					}
				}
			})
	}

}

func BenchmarkSolver_Solve(b *testing.B) {
	solver := NewSolver()
	for length := 10; length <= 20; length += 2 {
		b.Run(fmt.Sprintf("scarmble length %d", length),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					scrambled := cube.GetScrambled(20)
					_, err := solver.Solve(scrambled)
					if err != nil {
						b.Fatal(err.Error())
					}
				}
			})
	}
}
