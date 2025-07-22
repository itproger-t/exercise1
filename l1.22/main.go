package main

import (
	"fmt"
	"math/big"
)

func main() {
	a1 := int(1 << 21)
	b1 := int(1 << 22)

	fmt.Println("-- Операции с типом int --")
	fmt.Printf("a1 = %d, b1 = %d\n", a1, b1)
	fmt.Printf("a1 + b1 = %d\n", a1+b1)
	fmt.Printf("a1 - b1 = %d\n", a1-b1)
	fmt.Printf("a1 * b1 = %d\n", a1*b1)
	fmt.Printf("b1 / a1 = %d\n", b1/a1)

	a2 := big.NewInt(0).Lsh(big.NewInt(1), 100)
	b2 := big.NewInt(0).Lsh(big.NewInt(1), 120)

	fmt.Println("-- Операции с big.Int --")
	fmt.Printf("a2 = 2^100 = %s\n", a2.String())
	fmt.Printf("b2 = 2^120 = %s\n", b2.String())

	sum := big.NewInt(0).Add(a2, b2)
	fmt.Printf("a2 + b2 = %s\n", sum.String())

	diff := big.NewInt(0).Sub(b2, a2)
	fmt.Printf("b2 - a2 = %s\n", diff.String())

	product := big.NewInt(0).Mul(a2, b2)
	fmt.Printf("a2 * b2 = %s\n", product.String())

	quotient := big.NewInt(0).Div(b2, a2)
	fmt.Printf("b2 / a2 = %s\n", quotient.String())

}
