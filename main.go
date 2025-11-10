package main

import "fmt"

const (
	vatRate               = 0.20
	freeShippingThreshold = 3000.0
	shippingCost          = 300.0
)

func main() {
	product1Name := "Ноутбук"
	product1Price := 45000.0
	product1Qty := 1

	product2Name := "Мышь"
	product2Price := 8000.0
	product2Qty := 6

	product3Name := "Клавиатура"
	product3Price := 5000.0
	product3Qty := 1

	totalSumProduct := (product1Price * float64(product1Qty)) +
		(product2Price * float64(product2Qty)) +
		(product3Price * float64(product3Qty))

	totalVATRate := totalSumProduct * vatRate
	totalSum := totalSumProduct + totalVATRate

	var delivery string
	if totalSumProduct < freeShippingThreshold {
		totalSum += shippingCost
		delivery = "ПЛАТНО"
	} else {
		delivery = "БЕСПЛАТНО"
	}

	fmt.Println("╔════════════════════════════════════════════════╗")
	fmt.Println("║                    ЧЕК                         ║")
	fmt.Println("╚════════════════════════════════════════════════╝")
	fmt.Println("Товары")
	fmt.Printf(" 1. %s x%d %.2f\n", product1Name, product1Qty, product1Price)
	fmt.Printf(" 2. %s x%d %.2f\n", product2Name, product2Qty, product2Price)
	fmt.Printf(" 3. %s x%d %.2f\n", product3Name, product3Qty, product3Price)
	fmt.Printf("Сумма без НДС: %.2f\n", totalSumProduct)
	fmt.Printf("НДС(20): %.2f\n", totalVATRate)
	fmt.Printf("Доставка: %s\n", delivery)
	fmt.Printf("ИТОГО: %.2f", totalSum)
}
