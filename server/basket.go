package main

import "pokeStore/constants"

type ItemType string

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

var prices = map[ItemType]float32{
	REPELENTE:    5,
	CARAMELORARO: 20,
	SPORTS:       7.5,
}

func GetItemPrice(it ItemType) float32 {
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
	Items map[ItemType]int
}

func NewBasket(newId string) *Basket {
	return &Basket{
		Id:    newId,
		Items: make(map[ItemType]int),
	}
}

func (b *Basket) AddItem(it ItemType, amount int) {
	b.Items[it] += amount
}

func (b *Basket) GetItems() map[ItemType]int {
	return b.Items
}

func (b *Basket) RemoveItem(it ItemType, amount int) {
	b.Items[it] -= amount
	if b.Items[it] < 1 {
		b.Items[it] = 0
	}
}

func (b *Basket) Remove() {}
