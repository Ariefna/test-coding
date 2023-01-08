package main

import "fmt"

func barcharts(arraybar []int) {
	max := arraybar[0]
	for i := 1; i < len(arraybar); i++ {
		if max < arraybar[i] {
			max = arraybar[i]
		}
	}
	for x := max; x >= 1; x-- {
		for y := 1; y <= len(arraybar); y++ {
			if arraybar[y-1] >= x {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	for y := 1; y <= len(arraybar); y++ {
		fmt.Print(arraybar[y-1], " ")
	}
	fmt.Println()
}
func swap(arrayzor []int, i, j int) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}
func insertionSortdesc(array []int) {
	for i := len(array) - 1; i > 0; i-- {
		for j := i; j < len(array) && array[j] > array[j-1]; j++ {
			swap(array, j, j-1)
			barcharts(array)
		}
	}
}
func insertionSortasc(array []int) {
	for i := len(array) - 1; i > 0; i-- {
		for j := i; j < len(array) && array[j] < array[j-1]; j++ {
			swap(array, j, j-1)
			barcharts(array)
		}
	}
}
func main() {
	array := []int{1, 4, 5, 6, 8, 2}
	fmt.Println("- Vertical Barcharts")
	barcharts(array)
	fmt.Println("- Sorted array (ascending)")
	fmt.Println("- Steps visualization")
	insertionSortasc(array)
	fmt.Println("- Sorted array (descending)")
	fmt.Println("- Steps visualization")
	insertionSortdesc(array)
}
