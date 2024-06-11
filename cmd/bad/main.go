package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	Max   int
	Min   int = int(^uint(0) >> 1) // максимально можливе значення для int
	Sum   int
	Count int
	Med   int
	Avg   int
	Inca  []int
	Dica  []int
)

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left - 1
	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	pivotIndex := partition(nums, left, right)
	if k == pivotIndex {
		return nums[k]
	} else if k < pivotIndex {
		return quickSelect(nums, left, pivotIndex-1, k)
	} else {
		return quickSelect(nums, pivotIndex+1, right, k)
	}
}

func main() {
	start := time.Now()

	file, err := os.Open("files/10m.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var nums []int
	var incaT []int
	var dicaT []int
	lastSet := false
	var last int

	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Error converting line to int: %v\n", err)
			continue
		}
		nums = append(nums, n)
		if n > Max {
			Max = n
		}
		if n < Min {
			Min = n
		}
		Sum += n
		Count++

		if lastSet {
			if n > last {
				incaT = append(incaT, n)
			} else {
				if len(incaT) > len(Inca) {
					Inca = incaT
				}
				incaT = []int{n}
			}

			if n < last {
				dicaT = append(dicaT, n)
			} else {
				if len(dicaT) > len(Dica) {
					Dica = dicaT
				}
				dicaT = []int{n}
			}
		} else {
			incaT = append(incaT, n)
			dicaT = append(dicaT, n)
		}

		last = n
		lastSet = true
	}

	if len(incaT) > len(Inca) {
		Inca = incaT
	}
	if len(dicaT) > len(Dica) {
		Dica = dicaT
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	Avg = Sum / Count

	if Count%2 == 0 {
		Med = (quickSelect(nums, 0, len(nums)-1, Count/2-1) + quickSelect(nums, 0, len(nums)-1, Count/2)) / 2
	} else {
		Med = quickSelect(nums, 0, len(nums)-1, Count/2)
	}

	fmt.Printf("Максимальне число: %d\n", Max)
	fmt.Printf("Мінімальне число: %d\n", Min)
	fmt.Printf("Медіана: %d\n", Med)
	fmt.Printf("Середнє арифметичне значення: %d\n", Avg)
	fmt.Printf("Найбільша послідовність, яка збільшується: %v\n", Inca)
	fmt.Printf("Найбільша послідовність, яка зменшується: %v\n", Dica)

	log.Printf("It took %s", time.Since(start))
}
