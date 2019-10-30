package main

import "pokeStore/constants"

type ItemType string

type ItemPrice float32

const (
	REPELENTE    ItemType = constants.REPELENTE
	CARAMELORARO          = constants.CARAMELORARO
	SPORTS                = constants.SPORTS
)

var itemTypeRegistry = map[string]ItemType{
	constants.REPELENTE:    REPELENTE,
	constants.CARAMELORARO: CARAMELORARO,
	constants.SPORTS:       SPORTS,
}

var prices = map[ItemType]ItemPrice{
	REPELENTE:    5,
	CARAMELORARO: 20,
	SPORTS:       7.5,
}

func GetItemPrice(it ItemType) ItemPrice {
	if price, ok := prices[it]; ok {
		return price
	}
	return 0
}

func GetItemByType(raw string) (ItemType, bool) {
	itemType, ok := itemTypeRegistry[raw]
	return itemType, ok
}

// TODO: Provide better abstraction over basket id
//type BasketId string

type Basket struct {
	Id    string
	items map[ItemType]int
}

func NewBasket(newId string) *Basket {
	return &Basket{
		Id:    newId,
		items: make(map[ItemType]int),
	}
}

func (b *Basket) AddItem(it ItemType, amount int) {
	b.items[it] += amount
}

func (b *Basket) GetItemCount(it ItemType) int {
	if count, ok := b.items[it]; ok {
		return count
	}
	return 0
}

func (b *Basket) GetTotal() float32 {
	var result float32
	for itemType, itemCount := range b.items {
		result += float32(itemCount) * float32(GetItemPrice(itemType))
	}
	return result
}

func (b *Basket) RemoveItem(it ItemType, amount int) {
	b.items[it] -= amount
	if b.items[it] < 1 {
		b.items[it] = 0
	}
}
