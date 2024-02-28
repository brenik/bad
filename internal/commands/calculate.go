package commands

import (
	"fmt"
	"github/brenik/bad/internal/service"
	"github/brenik/bad/internal/storage"
	"log"
)

func Run(file string) {
	fmt.Println("File: ", file)
	lines, err := service.ReadFile(file)
	if err != nil {
		log.Fatalf("readFiles: %s", err)
	}
	storage.Lines = lines
	Compare()
	CalcMediana()
	CalcAverage()
	service.PrintResult()
}
