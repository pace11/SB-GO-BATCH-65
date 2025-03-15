package main

import (
	"fmt"
	"math"
	"strings"
)

var dataFilm = []map[string]string{}

func main() {

	// soal 1
	soalOne("Pace", "laki-laki", "Programmer", "29")

	// soal 2
	var fruits = []string{"semangka", "jeruk", "melon", "pepaya"}
	soalTwo("pace", fruits...)

	// soal 3
	soalThree()

	// soal 4 (faktorial)
	soalFour(3)

	// soal 5
	var luasLingkaran float64
	var kelilingLingkaran float64
	radius := 7.0
	hitungLingkarang(radius, &luasLingkaran, &kelilingLingkaran)
	fmt.Println("---- SOAL 5 ----")
	fmt.Println("----------------")
	fmt.Printf("Luas Lingkaran %.2f\n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran %.2f\n", kelilingLingkaran)
	fmt.Println()

}

func soalOne(name, gender, job, age string) {

	var aka string
	if gender == "laki-laki" {
		aka = "Pak"
	} else {
		aka = "Bu"
	}

	fmt.Println("---- SOAL 1 ----")
	fmt.Println("----------------")
	fmt.Printf("%s %s adalah seorang %s yang berusia %s tahun", aka, name, gender, job)
	fmt.Println()
	fmt.Println()
}

func soalTwo(name string, fruits ...string) {
	fmt.Println("---- SOAL 2 ----")
	fmt.Println("----------------")
	fmt.Printf(`Halo, nama saya %s dan buah favorit saya adalah "%s"`,
		name, strings.Join(fruits, `", "`))
	fmt.Println()
	fmt.Println()
}

func faktorial(number int) int {

	if number == 0 {
		return 1
	}
	return number * faktorial(number-1)
}

func tambahDataFilm() func(title, duration, genre, year string) {
	return func(title, duration, genre, year string) {
		film := map[string]string{
			"title":    title,
			"duration": duration,
			"genre":    genre,
			"year":     year,
		}
		dataFilm = append(dataFilm, film)
	}
}

func soalThree() {
	addFilm := tambahDataFilm()

	addFilm("LOTR", "2 jam", "action", "1999")
	addFilm("Avenger", "2 jam", "action", "2019")
	addFilm("Spiderman", "2 jam", "action", "2004")
	addFilm("Juon", "2 jam", "horror", "2004")

	fmt.Println("---- SOAL 3 ----")
	fmt.Println("----------------")
	for _, item := range dataFilm {
		fmt.Println(item)
	}
	fmt.Println()
}

func soalFour(number int) {
	fmt.Println("---- SOAL 4 ----")
	fmt.Println("----------------")
	fmt.Printf("Hasil faktorial dari %d: %d\n", number, faktorial(number))
	fmt.Println()
}

func hitungLingkarang(radius float64, luas *float64, keliling *float64) {
	*luas = math.Pi * radius * radius
	*keliling = 2 * math.Pi * radius
}
