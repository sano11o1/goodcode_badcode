package main

type Money struct {
	Value int
}

func (m Money) Add(n Money) Money {
	return Money{Value: m.Value + n.Value}
}

type HotelRate interface {
	Fee() Money
	BusySeasonFee() Money //繁忙期の料金もインターフェースで定義
}

type RegularRate struct{}

func (r RegularRate) Fee() Money {
	return Money{Value: 10000}
}

func (r RegularRate) BusySeasonFee() Money {
	return r.Fee().Add(Money{Value: 10000})
}

type PremiumRate struct{}

func (r PremiumRate) Fee() Money {
	return Money{Value: 20000}
}

func (r PremiumRate) BusySeasonFee() Money {
	return r.Fee().Add(Money{Value: 15000})
}
