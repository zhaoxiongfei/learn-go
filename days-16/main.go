package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %f", e)
}

func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, ErrNegativeSqrt(n)
	}
	z := 1.0
	for i := 0; i < 20; i++ {
		z = z - (z*z-n)/(2*n)
	}
	return z, nil
}

func main() {
	a, err := Sqrt(-10)
	if err != nil {
		fmt.Println("Sqrt error", err)
	} else {
		fmt.Println("Sqrt value:", -10, a)
	}
	b, err := Sqrt(10)
	if err != nil {
		fmt.Println("Sqrt error", err)
	} else {
		fmt.Println("Sqrt value:", 10, b)
	}

}
