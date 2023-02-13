package calculator

import (
	"github.com/mjmhtjain/calc_lib/extra"
	"github.com/mjmhtjain/calc_lib/internal/operations"
)

func Add(a, b int) int {
	return operations.Add(a, b)
}

func Sub(a, b int) int {
	return operations.Sub(a, b)
}

func Mul(a, b int) int {
	return operations.Mul(a, b)
}

func Div(a, b int) int {
	return extra.Div(a, b)
}
