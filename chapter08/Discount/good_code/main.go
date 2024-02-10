package main

func main() {}

type RegularPrice struct {
	Amount int
}

const MIN_AMOUNT = 10000

func NewRegularPrice(amount int) RegularPrice {
	if amount < MIN_AMOUNT {
		return RegularPrice{Amount: MIN_AMOUNT}
	}
	return RegularPrice{Amount: amount}
}

const REGULAR_DISCOUNT_MIN_AMOUNT = 10000
const REGULAR_DISCOUNT_AMOUNT = 400

type RegularDiscountedPrice struct {
	Amount int
}

func NewRegularDiscountedPrice(regularPrice RegularPrice) RegularDiscountedPrice {
	discountedAmount := regularPrice.Amount - REGULAR_DISCOUNT_AMOUNT
	if discountedAmount < REGULAR_DISCOUNT_MIN_AMOUNT {
		discountedAmount = REGULAR_DISCOUNT_MIN_AMOUNT
	}
	return RegularDiscountedPrice{Amount: discountedAmount}
}

const SUMMER_DISCOUNT_MIN_AMOUNT = 10000
const SUMMER_DISCOUNT_AMOUNT = 300

type SummerDiscountedPrice struct {
	Amount int
}

func NewSummerDiscountedPrice(regularPrice RegularPrice) RegularDiscountedPrice {
	discountedAmount := regularPrice.Amount - SUMMER_DISCOUNT_AMOUNT
	if discountedAmount < SUMMER_DISCOUNT_MIN_AMOUNT {
		discountedAmount = SUMMER_DISCOUNT_MIN_AMOUNT
	}
	return RegularDiscountedPrice{Amount: discountedAmount}
}
