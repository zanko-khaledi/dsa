package main

import (
	"fmt"
	"zanko-khaledi/dsa/algorithmes"
)

func main() {

	str := algorithmes.RandomStr(32)

	fmt.Println(str)

	fmt.Println(algorithmes.Frequency(str))

	fmt.Println(algorithmes.MostRepeatedChar(str))
}
