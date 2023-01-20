package main

import (
	"fmt"
	"strconv"
)

func main() {
	//i := uint16(2)
	i := uint64(2)
	var it interface{}
	it = i
	sv := fmt.Sprintf("%v", it)
	j, err := strconv.Atoi(sv)
	fmt.Println(err, j)

	//k, err := strconv.ParseFloat(sv, 64)
	k, err := strconv.ParseFloat("1.23", 64)
	fmt.Println(err, k)
}
