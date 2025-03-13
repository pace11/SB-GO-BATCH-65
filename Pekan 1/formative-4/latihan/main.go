package main

import (
	"fmt"
)

func main() {
	price := 10000
	category := "NEWUSER"
	distance := 1

	totalPayment := calculateTotalPayment(price, category, distance)

	fmt.Println("---------------------------------------------")
	fmt.Println("Bill Receipt ...")
	fmt.Println("---------------------------------------------")
	fmt.Printf("Price\t\t\t: Rp %d\n", price)
	fmt.Printf("Price After Discount\t: Rp %d\n", calculateDiscount(price, category))
	fmt.Printf("Tax\t\t\t: Rp %d\n", calculateTax(calculateDiscount(price, category)))
	fmt.Printf("Shipping Cost\t\t: Rp %d\n", CalculateShippingCost(distance))
	fmt.Printf("Total pembayaran\t: Rp %d\n", totalPayment)
	fmt.Println("---------------------------------------------")
}

func calculateDiscount(price int, category string) int {
	switch category {
	case "NEWUSER":
		return price * 80 / 100
	case "PROMO11":
		return price * 90 / 100
	default:
		return price
	}
}

func calculateTax(priceAfterDiscount int) int {
	return priceAfterDiscount * 10 / 100
}

func CalculateShippingCost(distance int) int {
	switch {
	case distance < 5:
		return 10000
	case distance >= 5 && distance <= 20:
		return 20000
	default:
		return 30000
	}
}

func calculateTotalPayment(price int, category string, distance int) int {
	priceAfterDiscount := calculateDiscount(price, category)
	tax := calculateTax(priceAfterDiscount)
	shippingCost := CalculateShippingCost(distance)
	total := priceAfterDiscount + tax + shippingCost
	return total
}
