package main

import (
	"fmt"
	"golang-dasar/database"
)

/*
Blank Identifier

Jika tidak ingin menjalankan init function tanpa harus mengeksekusinya salah satu
function yang ada di package.

Maka cara menanganinya adalah menggunakan blank identifier ( garis bawah ( _ ) )
sebelum nama package ketika melakukan import. Contohnya:
_ "golang-dasar/database"
*/

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
