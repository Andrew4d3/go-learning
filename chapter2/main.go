package main

import "fmt"

const value = 42

func main() {
	var i int = 20
	var f float32 = float32(i)

	fmt.Println(i, f)

	i = value
	f = value
	fmt.Println(i, f)

	var b int8 = 127
	var smalli int32 = 2147483647
	var bigi int64 = 9223372036854775807

	fmt.Println(b+1, smalli+1, bigi+1)
}
