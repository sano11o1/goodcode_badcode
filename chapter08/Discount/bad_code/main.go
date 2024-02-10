package main

import (
	"errors"
	"fmt"
)

func main() {
	// 通常の割引
	manager := NewDiscountManager(
		[]ProductDiscount{
			{ID: "1", CanDiscount: true},
		},
	)
	product := Product{
		ID:    "1",
		Name:  "apple",
		Price: 1000,
	}
	err := manager.Add(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(manager.TotalPrice) // 700

	// 夏季割引
	manager2 := NewDiscountManager(
		[]ProductDiscount{
			{ID: "1", CanDiscount: true},
		},
	)
	product2 := Product{
		ID:          "1",
		Name:        "apple",
		Price:       1000,
		CanDiscount: true,
	}
	summerManager := NewSummerDiscountManager(manager2)
	err = summerManager.Add(product2)
	if err != nil {
		panic(err)
	}
	fmt.Println(summerManager.DiscountManager.TotalPrice) // 700（getDiscountPriceの実装と連動する)
}

type Product struct {
	ID          string
	Name        string
	Price       int
	CanDiscount bool // 新規追加 「夏季」割引可能な場合はtrue
}

type ProductDiscount struct {
	ID          string
	CanDiscount bool
}

type DiscountManager struct {
	TotalPrice       int
	DiscountProducts []Product
	ProductDiscounts []ProductDiscount
}

func NewDiscountManager(productDiscounts []ProductDiscount) DiscountManager {
	return DiscountManager{
		ProductDiscounts: productDiscounts,
	}
}

/*
商品1点につき300円引き
上限20000円まで商品を追加可能
追加に成功した場合はnilを返す
*/
func (d *DiscountManager) Add(product Product) error {
	if product.ID == "" {
		return errors.New("Invalid product ID")
	}
	if product.Name == "" {
		return errors.New("Invalid product name")
	}
	if product.Price < 0 {
		return errors.New("Invalid product price")
	}

	var discountProduct ProductDiscount
	for _, p := range d.ProductDiscounts {
		if p.ID == product.ID {
			discountProduct = p
			break
		}
	}

	if product.ID != discountProduct.ID {
		return errors.New("Invalid product ID")
	}

	discountPrice := d.getDiscountPrice(product.Price)

	var tmp int
	if discountProduct.CanDiscount {
		tmp = d.TotalPrice + discountPrice
	} else {
		tmp = d.TotalPrice + product.Price
	}

	if tmp <= 20000 {
		d.TotalPrice = tmp
		d.DiscountProducts = append(d.DiscountProducts, product)
	} else {
		return errors.New("Total price is over 20000")
	}

	return nil
}

func (d *DiscountManager) getDiscountPrice(price int) int {
	discountPrice := price - 300 // ここの300円を変更するとSummerDiscountManagerの割引額も変更されてしまう
	if discountPrice < 0 {
		return 0
	}
	return discountPrice
}

type SummerDiscountManager struct {
	DiscountManager DiscountManager
}

func NewSummerDiscountManager(manager DiscountManager) SummerDiscountManager {
	return SummerDiscountManager{
		DiscountManager: manager,
	}
}

/*
商品1点につき300円引き
上限30000円まで商品を追加可能
*/
func (s *SummerDiscountManager) Add(product Product) error {
	if product.ID == "" {
		return errors.New("Invalid product ID")
	}
	if product.Name == "" {
		return errors.New("Invalid product name")
	}

	var tmp int
	if product.CanDiscount {
		tmp = s.DiscountManager.TotalPrice + s.DiscountManager.getDiscountPrice(product.Price)
	} else {
		tmp = s.DiscountManager.TotalPrice + product.Price
	}

	if tmp < 30000 {
		s.DiscountManager.TotalPrice = tmp
		s.DiscountManager.DiscountProducts = append(s.DiscountManager.DiscountProducts, product)
	} else {
		return errors.New("Total price is over 30000")
	}

	return nil
}
