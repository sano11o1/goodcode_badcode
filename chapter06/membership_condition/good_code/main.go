package main

func main() {
	history := History{
		TotalAmount:               100000,
		PurchaseFrequencyPerMonth: 10,
		ReturnRate:                0.0001,
	}
	policy := NewGoldCustomerPolicy()
	if policy.complyWithAll(history) {
		println("Gold Customer")
	}

	history = History{
		TotalAmount:               1000,
		PurchaseFrequencyPerMonth: 10,
		ReturnRate:                0.0001,
	}
	if policy.complyWithAll(history) {
		println("Gold Customer")
	}
}

type History struct {
	TotalAmount               int
	PurchaseFrequencyPerMonth int
	ReturnRate                float64
}

type ExcellentCustomerRule interface {
	ok(History) bool
}

type GoldCustomerPurchaseAmountRule struct{}

func (r GoldCustomerPurchaseAmountRule) ok(history History) bool {
	return 100000 <= history.TotalAmount
}

type CustomerPurchaseFrequencyRule struct{}

func (r CustomerPurchaseFrequencyRule) ok(history History) bool {
	return 10 <= history.PurchaseFrequencyPerMonth
}

type CustomerReturnRateRule struct{}

func (r CustomerReturnRateRule) ok(history History) bool {
	return history.ReturnRate <= 0.0001
}

type ExcellentCustomerPolicy struct {
	Rules []ExcellentCustomerRule
}

func (p *ExcellentCustomerPolicy) add(rule ExcellentCustomerRule) {
	p.Rules = append(p.Rules, rule)
}

func (p ExcellentCustomerPolicy) complyWithAll(history History) bool {
	for _, rule := range p.Rules {
		if !rule.ok(history) {
			return false
		}
	}
	return true
}

type GoldCustomerPolicy struct {
	ExcellentCustomerPolicy
}

func (p GoldCustomerPolicy) complyWithAll(history History) bool {
	return p.ExcellentCustomerPolicy.complyWithAll(history)
}

func NewGoldCustomerPolicy() GoldCustomerPolicy {
	policy := &ExcellentCustomerPolicy{}
	policy.add(GoldCustomerPurchaseAmountRule{})
	policy.add(CustomerPurchaseFrequencyRule{})
	policy.add(CustomerReturnRateRule{})
	return GoldCustomerPolicy{ExcellentCustomerPolicy: *policy}
}
