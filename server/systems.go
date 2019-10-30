package main

type DiscountSystem interface {
	applyDiscount(Basket) float32
}

type MarketingDiscountSystem struct {
	/*
		The marketing department thinks a buy 2 get 1 free promotion will work
		best (buy two of the same product, get one free), and would like this
		to only apply to REPELENTE items.
	*/
	targetType ItemType
}

func NewMarketingDiscountSystem(tt ItemType) *MarketingDiscountSystem {
	return &MarketingDiscountSystem{targetType: tt}
}

func (m *MarketingDiscountSystem) applyDiscount(basket Basket) float32 {
	allItemsCount := basket.GetItems()
	if itemCount, ok := allItemsCount[m.targetType]; ok {
		if itemCount < 2 {
			return 0
		}
		return float32(itemCount/2) * GetItemPrice(m.targetType)
	}
	return 0
}

type CFODiscountSystem struct {
	/*
		The CFO insists that the best way to increase sales is with discounts on
		bulk purchases (buying x or more of a product, the price of that product
		is reduced), and requests that if you buy 3 or more CARAMELORARO items, the
		price per unit should be 19.00€
	*/
	targetType      ItemType
	priceIfDiscount float32
}

func NewCFODiscountSystem(tt ItemType, discount float32) *CFODiscountSystem {
	return &CFODiscountSystem{targetType: tt, priceIfDiscount: discount}
}

func (s *CFODiscountSystem) applyDiscount(basket Basket) float32 {
	allItemsCount := basket.GetItems()
	if itemCount, ok := allItemsCount[s.targetType]; ok {
		if itemCount < 3 {
			return 0.0
		}
		itemPrice := GetItemPrice(s.targetType)
		return (itemPrice - s.priceIfDiscount) * float32(itemCount)
	}
	return 0.0
}

type CheckOutSystem struct {
	discountSystems []DiscountSystem
}

func (c *CheckOutSystem) RegisterDiscount(d DiscountSystem) {
	c.discountSystems = append(c.discountSystems, d)
}

func (c *CheckOutSystem) CheckOut(basket Basket) (float32, float32) {
	var total float32 = 0.0

	for itemType, itemCount := range basket.GetItems() {
		total += GetItemPrice(itemType) * float32(itemCount)
	}

	var totalDiscount float32 = 0.0

	for _, system := range c.discountSystems {
		totalDiscount += system.applyDiscount(basket)
	}

	return total, totalDiscount
}
