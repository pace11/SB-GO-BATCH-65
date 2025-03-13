// operator & conditional

package main

import "fmt"

func main() {
	// penjumlahan
	var add = 2 + 3
	fmt.Println(add)

	// pengurangan
	var subs = add - 3
	fmt.Println(subs)

	// pembagian
	var div = add / 3
	fmt.Println(div)

	// perkalian
	var multi = add * 2
	fmt.Println(multi)

	// modulus
	modulus := 8 % 3
	fmt.Println(modulus)

	// operasi matematika
	var angka = 8
	fmt.Println(angka)
	angka += 10
	fmt.Println(angka)

	var angka2 = 5
	fmt.Println(angka2)
	angka2 += 5
	fmt.Println(angka2)

	// operator perbandingan
	var angkaPembanding = 8

	fmt.Println(angkaPembanding > 5)

	fmt.Println(angkaPembanding < 5)

	fmt.Println(angkaPembanding >= 5)

	fmt.Println(angkaPembanding <= 5)

	fmt.Println(angkaPembanding == 5)

	fmt.Println(angkaPembanding != 5)

	// operator logika
	var a = true
	var b = false
	var c = true
	var d = false

	fmt.Println(a && c)

	fmt.Println(a && b)

	fmt.Println(a || b)

	fmt.Println(b || d)

	fmt.Println(!b && !d)

	fmt.Println(!a || b)

	// kondisi IF ELSE
	var minimarketStatus = "close"
	var minuteRemainingToOpen = 5
	if minimarketStatus == "open" {
		fmt.Println("saya akan membeli telur dan buah")
	} else if minuteRemainingToOpen <= 5 {
		fmt.Println("minimarket buka sebentar lagi, saya tungguin")
	} else {
		fmt.Println("minimarket tutup, saya pulang lagi")
	}

	// variable temporary di IF ELSE
	if minimarketStatus, minuteRemainingToOpen := "close", 5; minimarketStatus == "open" {
		fmt.Println("saya akan membeli telur dan buah")
	} else if minuteRemainingToOpen <= 5 {
		fmt.Println("minimarket buka sebentar lagi, saya tungguin")
	} else {
		fmt.Println("minimarket tutup, saya pulang lagi")
	}

	// kondisi SWITCH CASE
	var buttonPushed = 1
	switch buttonPushed {
	case 1:
		fmt.Println("matikan TV!")
	case 2, 3, 4:
		fmt.Println("turunkan volume TV!")
	default:
		fmt.Println("Tidak terjadi apa-apa")
	}

	// kondisi SWITCH CASE dengan kondisi seperti IF ELSE
	var buttonPushedWithIfElse = 1
	switch {
	case buttonPushedWithIfElse == 1:
		fmt.Println("matikan TV!")
	case buttonPushedWithIfElse == 2 && buttonPushedWithIfElse == 3:
		fmt.Println("turunkan volume TV!")
	case buttonPushedWithIfElse == 4:
		fmt.Println("matikan suara TV!")
	default:
		fmt.Println("Tidak terjadi apa-apa")
	}

	// kondisi SWITCH CASE dengan fallthrough
	var point = 6
	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}
