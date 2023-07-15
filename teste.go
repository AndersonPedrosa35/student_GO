import (
	"fmt"
)

// interface

func sum(a, b int) int {
	return a + b
}

func subtr(a, b int) (int, err) {
	if a < b {
		return 0, fmt.Errorf("O 1 parametro é maior que o segundo")
	}

	return a - b, nil
}

func mult(a, b int) (int, err) {
	if a > 1  b > 1 {
		return 0, fmt
	}
}

