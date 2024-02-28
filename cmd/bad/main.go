package main

import (
	"fmt"
	"github/brenik/bad/internal/commands"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello Porta")
	start := time.Now()
	file := "files/10m.txt"
	commands.Run(file)
	log.Printf("It took %s", time.Since(start))
}
