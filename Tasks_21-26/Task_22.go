/*

№ 22 Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0)
	b := big.NewInt(0)
	a.SetString("978697859596978878596586878595678", 10)
	b.SetString("223342515343435254445521543352434", 10)

	fmt.Println(addition(a, b))
	fmt.Println(subtraction(a, b))
	fmt.Println(multiplication(a, b))
	fmt.Println(division(a, b))
}

func addition(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func subtraction(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func multiplication(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func division(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}
