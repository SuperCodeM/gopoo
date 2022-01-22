package course

import "fmt"

type Course struct {
	Name    string
	Price   float64
	isFree  bool
	userIDs []uint
	clases  map[uint]string
}

func (c *Course) Changeprice(price float64){
	c.Price = price
}

func (c *Course) printclases (){
	text := "Las clases que tenemos son: "
	for _, class := range c.clases{
    text += class + ", "
	}
	fmt.Println(text[:len (text)-2])
}