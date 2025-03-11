package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// --> contoh dengan var <--
	var text = "Hello World"
	fmt.Println(text)

	text = "Hello World diubah"
	fmt.Println(text)

	// --> contoh dengan var dan tipe data <--
	var text1 string
	text1 = "ini teks 1"
	fmt.Println(text1)

	var text2 string = "ini teks 2"
	text2 = "ini teks 1 diubah"
	fmt.Println(text2)

	// --> contoh dengan perantara := <--
	text3 := "ini teks 1"
	text3 = "ini teks 1 diubah"
	fmt.Println(text3)

	text4 := "ini teks 2"
	fmt.Println(text4)

	// --> contoh tipe data string <--
	var name string = "Pace"
	fmt.Println(name)

	// --> contoh tipe data int <--
	var age int = 29
	fmt.Println(age)

	//  --> contoh tipe data int <--
	var floatNumber float32 = 2.9
	fmt.Println(floatNumber)

	// --> contoh tipe data boolean <--
	var isUser bool
	fmt.Println(isUser)

	isUser = true
	fmt.Println(isUser)

	// --> konversi <--

	var angkaFloat float64 = float64(age)
	var angkaInt int = int(floatNumber)
	fmt.Println(angkaFloat)
	fmt.Println(angkaInt)

	// --> mengambil huruf dari word berdasarkan index <--
	fmt.Println(string(name[0]))

	// --> replace kata tertentu dengan kata yang baru <--
	str := "tes test test"
	str = strings.Replace(str, "tes", "coba", 1)
	fmt.Println(str)

	// --> penjumlahan dengan menyamakan tipe data terlebih dahulu <--
	strAngka := "10"
	convAngka, _ := strconv.Atoi(strAngka)
	jumlah := 5 + convAngka
	fmt.Println(jumlah)

}
