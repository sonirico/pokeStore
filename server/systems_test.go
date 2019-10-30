package main

import (
	"testing"
)

type SystemTest struct {
	basket           *Basket
	expectedTotal    float32
	expectedDiscount float32
}

func testDiscountSystem(t *testing.T, tests []SystemTest, system *CheckOutSystem) bool {
	t.Helper()

	for _, test := range tests {
		actualTotal, actualDiscount := system.CheckOut(*test.basket)
		if test.expectedTotal != actualTotal {
			t.Errorf("Expected basket total to be %g. Got %g", test.expectedTotal,
				actualTotal)
			return false
		}
		if test.expectedDiscount != actualDiscount {
			t.Errorf("Expected basket discount to be %g. Got %g", test.expectedDiscount,
				actualDiscount)
			return false
		}
	}
	return true
}

func TestMarketingDiscountSystem(t *testing.T) {
	tests := []SystemTest{
		{
			basket: &Basket{
				items: map[ItemType]int{
					REPELENTE: 2,
				},
			},
			expectedTotal:    10,
			expectedDiscount: 5,
		},
		{
			basket: &Basket{
				items: map[ItemType]int{
					REPELENTE: 1,
				},
			},
			expectedTotal:    5,
			expectedDiscount: 0,
		},
		{
			basket: &Basket{
				items: map[ItemType]int{
					REPELENTE: 0,
				},
			},
			expectedTotal:    0,
			expectedDiscount: 0,
		},
		{
			basket: &Basket{
				items: map[ItemType]int{
					REPELENTE: 4,
				},
			},
			expectedTotal:    20,
			expectedDiscount: 10,
		},
		{
			basket: &Basket{
				items: map[ItemType]int{
					REPELENTE: 5,
				},
			},
			expectedTotal:    25,
			expectedDiscount: 10,
		},
		// Only affects REPELENTE items
		{
			basket: &Basket{
				items: map[ItemType]int{
					SPORTS: 10,
				},
			},
			expectedTotal:    75,
			expectedDiscount: 0,
		},
		// Only affects REPELENTE items
		{
			basket: &Basket{
				items: map[ItemType]int{
					CARAMELORARO: 10,
				},
			},
			expectedTotal:    200,
			expectedDiscount: 0,
		},
	}

	system := &CheckOutSystem{}
	system.RegisterDiscount(&MarketingDiscountSystem{targetType: REPELENTE})

	testDiscountSystem(t, tests, system)
}

func TestCFODiscountSystem(t *testing.T) {
	tests := []SystemTest{
		{
			basket: &Basket{
				items: map[ItemType]int{
					CARAMELORARO: 1,
				},
			},
			expectedTotal:    20,
			expectedDiscount: 0,
		},
		{
			basket: &Basket{
				items: map[ItemType]int{
					CARAMELORARO: 2,
				},
			},
			expectedTotal:    40,
			expectedDiscount: 0,
		},
		{
			basket: &Basket{
				items: map[ItemType]int{
					CARAMELORARO: 3,
				},
			},
			expectedTotal:    60,
			expectedDiscount: 3,
		},
		{
			basket: &Basket{
				items: map[ItemType]int{
					CARAMELORARO: 4,
				},
			},
			expectedTotal:    80,
			expectedDiscount: 4,
		},
		{
			basket: &Basket{
				items: map[ItemType]int{
					CARAMELORARO: 10,
				},
			},
			expectedTotal:    200,
			expectedDiscount: 10,
		},
		// Only affects CARAMELORAROS items
		{
			basket: &Basket{
				items: map[ItemType]int{
					SPORTS: 10,
				},
			},
			expectedTotal:    75,
			expectedDiscount: 0,
		},
		// Only affects CARAMELORAROS items
		{
			basket: &Basket{
				items: map[ItemType]int{
					REPELENTE: 10,
				},
			},
			expectedTotal:    50,
			expectedDiscount: 0,
		},
	}

	system := &CheckOutSystem{}
	system.RegisterDiscount(&CFODiscountSystem{targetType: CARAMELORARO, priceIfDiscount: 19.0})

	testDiscountSystem(t, tests, system)
}
