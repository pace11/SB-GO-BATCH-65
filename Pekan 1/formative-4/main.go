package main

import (
	"fmt"
)

func main() {
	// Soal 1
	soalOne(3)

	// soal 2
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	soalTwo(kalimat[:])

	// soal 3
	var vegetables = []string{
		"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun",
	}
	soalThree(vegetables)

	// soal 4
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	soalFour(satuan)

	// soal 5
	length := 12
	width := 4
	height := 8
	soalFive(length, width, height)

}

func soalOne(length int) {
	fmt.Println("---- SOAL 1 ----")
	fmt.Println("length: ", length)
	fmt.Println("----------------")
	for i := 1; i <= length; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("#")
		}
		fmt.Println()
	}
	fmt.Println()
}

func soalTwo(words []string) {
	fmt.Println("---- SOAL 2 ----")
	fmt.Println("----------------")
	fmt.Println(words[2:])
	fmt.Println()
}

func soalThree(vegetables []string) {
	fmt.Println("---- SOAL 3 ----")
	fmt.Println("----------------")

	tmp := []string{}

	// append dari list vegetables ke tmp
	tmp = append(tmp, vegetables...)

	for index, item := range tmp {
		fmt.Printf("%d. %s\n", index+1, item)
	}
	fmt.Println()
}

func soalFour(value map[string]int) {
	fmt.Println("---- SOAL 4 ----")
	fmt.Println("----------------")

	hasil := 1
	for key, value := range value {
		fmt.Printf("%s: %d\n", key, value)
		hasil *= value
	}

	fmt.Println("Hasil perkalian:", hasil)
	fmt.Println()
}

func soalFive(length int, width int, height int) {
	fmt.Println("---- SOAL 5 ----")
	fmt.Println("----------------")
	fmt.Println("Panjang\t:", length)
	fmt.Println("Lebar\t:", width)
	fmt.Println("Tinggi\t:", height)
	fmt.Println("----------------")
	fmt.Println("Luas\t:", luas(length, width))
	fmt.Println("Keliling:", keliling(length, width))
	fmt.Println("Volume\t:", volume(length, width, height))
}

func luas(length int, width int) int {
	return length * width
}

func keliling(length int, width int) int {
	return 2 * (length + width)
}

func volume(length int, width int, height int) int {
	return length * width * height
}
