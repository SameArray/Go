package main

import "fmt"

func main() {
	shoppingCart := ShoppingCart{}

	shoppingCart.SetPaymentStrategy(CreditCardPayment{})
	shoppingCart.Checkout(100)

	shoppingCart.SetPaymentStrategy(KaspiQRPayment{})
	shoppingCart.Checkout(200)
}

type PaymentStrategy interface {
	Pay(amount int)
}

type CreditCardPayment struct{}

func (c CreditCardPayment) Pay(amount int) {
	fmt.Println("Оплата кредитной картой на сумму", amount)
}

type KaspiQRPayment struct{}

func (p KaspiQRPayment) Pay(amount int) {
	fmt.Println("Оплата каспи QR на сумму: ", amount)
}

type ShoppingCart struct {
	PaymentStrategy PaymentStrategy
}

func (s *ShoppingCart) SetPaymentStrategy(ps PaymentStrategy) {
	s.PaymentStrategy = ps
}

func (s *ShoppingCart) Checkout(amount int) {
	s.PaymentStrategy.Pay(amount)
}
