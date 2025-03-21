package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

func tampilkanPesan(kalimat string, tahun int) {
	defer fmt.Println("Function ini dieksekusi terakhir!")
	fmt.Printf("Kalimat: %s\nTahun: %d\n", kalimat, tahun)
}

func kelilingSegitigaSamaSisi(sisi int, withText bool) (string, error) {
	if sisi == 0 {
		err := errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		if !withText {
			// Jika withText false, langsung panic
			panic(err)
		}
		return "", err
	}

	keliling := sisi * 3
	if withText {
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	}
	return fmt.Sprintf("%d", keliling), nil
}

func tambahAngka(n int, angka *int) {
	*angka += n
}

func cetakAngka(angka *int) {
	fmt.Println("Total angka:", *angka)
}

func tambahPhone(phones *[]string, phone string) {
	*phones = append(*phones, phone)
}

func tampilkanPhones(phones []string) {
	sort.Strings(phones)
	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}

func tampilkanPhonesFive(phones []string, wg *sync.WaitGroup) {
	defer wg.Done()

	sort.Strings(phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}

func soalOne() {
	fmt.Println("Program dimulai...")
	defer tampilkanPesan("Golang Backend Development", 2021)
	fmt.Println("Program selesai...")
}

func soalTwo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi panic:", r)
		}
	}()
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))
}

func soalThree() {
	angka := 1
	defer cetakAngka(&angka)
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)
}

func soalFour() {
	var phones = []string{}
	tambahPhone(&phones, "Xiaomi")
	tambahPhone(&phones, "Asus")
	tambahPhone(&phones, "IPhone")
	tambahPhone(&phones, "Samsung")
	tambahPhone(&phones, "Oppo")
	tambahPhone(&phones, "Realme")
	tambahPhone(&phones, "Vivo")
	tampilkanPhones(phones)
}

func soalFive() {
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	var wg sync.WaitGroup
	wg.Add(1)
	go tampilkanPhonesFive(phones, &wg)
	wg.Wait()
}

func main() {
	fmt.Println("---- SOAL 1 ----")
	soalOne()

	fmt.Println("\n---- SOAL 2 ----")
	soalTwo()

	fmt.Println("\n---- SOAL 3 ----")
	soalThree()

	fmt.Println("\n---- SOAL 4 ----")
	soalFour()

	fmt.Println("\n---- SOAL 5 ----")
	soalFive()
}
