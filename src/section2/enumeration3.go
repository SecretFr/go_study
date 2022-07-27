package main

import "fmt"

func main() {
	const (
		_ = iota
		A
		_ // skip 처리
		C
	)

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		GOLD
		PLATINUM
	)

	fmt.Println(DEFAULT)
	fmt.Println(SILVER)
	// fmt.Println(_)
	fmt.Println(GOLD)
	fmt.Println(PLATINUM)

	type ByteSize float64

	const (
		_           = iota // ignore first value by assigning to blank identifier
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(ZB)
	fmt.Println(YB)
}
