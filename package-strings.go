package main

import (
	"fmt"
	"strings"
)

// https://golang.org/pkg/

func main() {
	fmt.Println(strings.Contains("Albany Siswanto", "Albany")) // true
	fmt.Println(strings.Contains("Albany Siswanto", "Budi"))   // false

	fmt.Println(strings.Split("Albany Siswanto", " ")) // menjadi slice

	fmt.Println(strings.ToLower("Albany Siswanto")) // albany siswanto
	fmt.Println(strings.ToUpper("Albany Siswanto")) // ALBANY SISWANTO
	fmt.Println(strings.ToTitle("Albany Siswanto")) // ALBANY SISWANWTO

	fmt.Println(strings.Trim("  Albany Siswanto  ", " ")) // menghilangkan diawal dan diakhir

	fmt.Println(strings.ReplaceAll("Albany Siswanto", "Albany", "Semmy"))
}
