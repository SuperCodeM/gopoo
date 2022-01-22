package course

import "fmt"

type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Clases  map[uint]string
}

func new(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}

	return &course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}
}

func (c *course) Changeprice(price float64) {
	c.Price = price
}

func (c *course) Printclases() {
	text := "Las clases que tenemos son: "
	for _, class := range c.Clases {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}
