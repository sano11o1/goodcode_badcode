package main

func main() {}

type History struct {
	TotalAmount               int
	PurchaseFrequencyPerMonth int
	ReturnRate                float64
}

func isGoldCustomer(history History) bool {
	if 100000 <= history.TotalAmount {
		if 10 <= history.PurchaseFrequencyPerMonth {
			if history.ReturnRate <= 0.0001 {
				return true
			}
		}
	}
	return false
}

func isSilverCustomer(history History) bool {
	if 10 <= history.PurchaseFrequencyPerMonth {
		if history.ReturnRate <= 0.0001 {
			return true
		}
	}
	return false
}
