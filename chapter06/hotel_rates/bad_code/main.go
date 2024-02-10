package main

type Money struct {
	Value int
}

func (m Money) Add(n Money) Money {
	return Money{Value: m.Value + n.Value}
}

type HotelRate interface {
	Fee() Money
}

type RegularRate struct{}

func (r RegularRate) Fee() Money {
	return Money{Value: 10000}
}

type PremiumRate struct{}

func (r PremiumRate) Fee() Money {
	return Money{Value: 20000}
}

func main() {
	// 繁忙期でないときの料金
	regular := RegularRate{}
	premium := PremiumRate{}
	println(regular.Fee().Value)
	println(premium.Fee().Value)

	// 繁忙期の料金
	println(BusySeasonFee(regular).Value)
}

// せっかくInterfaceを使っているのに、関数内で型を判定している！！！
func BusySeasonFee(rate HotelRate) Money {
	if _, ok := rate.(PremiumRate); ok {
		return rate.Fee().Add(Money{Value: 10000})
	}
	if _, ok := rate.(RegularRate); ok {
		return rate.Fee().Add(Money{Value: 15000})
	}
	return Money{Value: 0}
}
