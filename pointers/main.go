package main

import (
	"fmt"
	"os"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 90
}

func main() {
	x := 10
	p := &x //p é um ponteiro para x
	y := *p //y é atribuido o valor apontado por p

	fmt.Println("Value of x:", x)
	fmt.Println("Memory address of x:", p)
	fmt.Println("Value of pointer address of x:", y)

	fmt.Fprintln(os.Stdout, []any{"\n"}...)

	i := 88
	fmt.Println("Initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer: ", &i)
}
