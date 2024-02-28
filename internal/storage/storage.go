package storage

import "fmt"

var Min = 0
var Max = 0
var Sum = 0
var Med = 0
var Avg = 0
var Inc = make(map[int]int)
var Dic = make(map[int]int)

var Lines []string

func Print(label string) {
	fmt.Println(label)
}
