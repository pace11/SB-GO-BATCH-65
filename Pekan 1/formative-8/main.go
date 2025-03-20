package main

import (
	"fmt"
	"formative-8/matematika"
	"math"
	"strings"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type phone struct {
	name   string
	brand  string
	year   int
	colors []string
}

type phoneInfo interface {
	display() string
}

func (p phone) display() string {
	return fmt.Sprintf(
		"name: %s\nbrand: %s\nyear: %d\ncolors: %s",
		p.name, p.brand, p.year, strings.Join(p.colors, ", "),
	)
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * float64((b.panjang*b.lebar)+(b.panjang*b.tinggi)+(b.lebar*b.tinggi))
}

func luasPersegi(sisi int, withText bool) interface{} {
	if sisi == 0 {
		if withText {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	luas := sisi * sisi
	if withText {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}
	return luas
}

func main() {
	segitiga := segitigaSamaSisi{alas: 10, tinggi: 8}
	persegi := persegiPanjang{panjang: 5, lebar: 3}
	tabung1 := tabung{jariJari: 7, tinggi: 10}
	balok1 := balok{panjang: 4, lebar: 3, tinggi: 5}

	fmt.Println("---- SOAL 1 ----")
	fmt.Println("---- Bangun Datar ----")
	fmt.Printf("Luas Segitiga Sama Sisi: %d\n", segitiga.luas())
	fmt.Printf("Keliling Segitiga Sama Sisi: %d\n", segitiga.keliling())
	fmt.Printf("Luas Persegi Panjang: %d\n", persegi.luas())
	fmt.Printf("Keliling Persegi Panjang: %d\n", persegi.keliling())

	fmt.Println("\n---- Bangun Ruang ----")
	fmt.Printf("Volume Tabung: %.2f\n", tabung1.volume())
	fmt.Printf("Luas Permukaan Tabung: %.2f\n", tabung1.luasPermukaan())
	fmt.Printf("Volume Balok: %.2f\n", balok1.volume())
	fmt.Printf("Luas Permukaan Balok: %.2f\n", balok1.luasPermukaan())

	myPhone := phone{
		name:   "Samsung",
		brand:  "Samsung Galaxy",
		year:   2020,
		colors: []string{"Bronze", "Silver", "Black"},
	}

	fmt.Println("\n---- SOAL 2 ----")
	fmt.Println(myPhone.display())

	fmt.Println("\n---- SOAL 3 ----")
	fmt.Println(luasPersegi(4, true))

	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// Type assertion untuk mengakses nilai dari interface{}
	prefixText := prefix.(string)
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	allNumbers := append(angkaPertama, angkaKedua...)

	total := 0
	numberStrings := []string{}
	for _, num := range allNumbers {
		total += num
		numberStrings = append(numberStrings, fmt.Sprintf("%d", num))
	}

	resultString := prefixText + strings.Join(numberStrings, " + ") + fmt.Sprintf(" = %d", total)
	fmt.Println("\n---- SOAL 4 ----")
	fmt.Println(resultString)

	a, b := 5, 3

	hasilTambah := matematika.Tambah(a, b)
	hasilKali := matematika.Kali(a, b)

	fmt.Println("\n---- SOAL 5 ----")
	fmt.Println("Hasil Penjumlahan:", hasilTambah) // Output: 8
	fmt.Println("Hasil Perkalian:", hasilKali)     // Output: 15
}
