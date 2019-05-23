package cart

import "fmt"

type Cart struct {
	items map[string]Item
}
type Item struct {
	ID    string
	Name  string
	Price float64
	Qty   int
}

// AddItem adds an item to the cart
func (c *Cart) AddItem(i Item) {
	c.init()
	if existingItem, ok := c.items[i.ID]; ok {
		existingItem.Qty++
		c.items[i.ID] = existingItem

	} else {
		i.Qty = 1
		c.items[i.ID] = i
	}
}

// RemoveItem removes n number of items with give id from the cart
func (c *Cart) RemoveItem(id string, n int) {
	c.init()
	if existingItem, ok := c.items[id]; ok {
		if existingItem.Qty <= n {
			delete(c.items, id)
		} else {
			existingItem.Qty -= n
			c.items[id] = existingItem
		}
	}
}

func (c *Cart) TotalAmount() float64 {
	c.init()
	totalAmount := 0.00
	for _, item := range c.items {
		totalAmount += float64(item.Qty) * item.Price
	}
	return totalAmount
}

func (c *Cart) TotalUnit() int {
	c.init()
	totalUnit := 0
	for _, i := range c.items {
		totalUnit += i.Qty
	}
	return totalUnit
}

func (c *Cart) TotalUniqueItems() int {

	return len(c.items)
}

func (c *Cart) init() {
	if c.items == nil {
		c.items = map[string]Item{}
	}
}

type CustomMap map[string]int

type CustomCart struct {
	items CustomMap
}

func main() {
	apple := Item{
		ID:    "001",
		Name:  "apple",
		Price: 3.00,
		Qty:   2,
	}

	grape := Item{
		ID:    "002",
		Name:  "grape",
		Price: 4.00,
		Qty:   1,
	}

	items1 := map[string]Item{
		"items1": apple,
		"items2": grape,
	}
	cart1 := Cart{
		items: map[string]Item{
			"item1": apple,
			"item2": grape,
		},
	}

	customMap1 := CustomMap{
		"item3": 10,
		"item4": 20,
	}
	for item, num := range items1 {
		fmt.Println("item is ", item)
		fmt.Println("num is ", num)

	}

	fmt.Println(cart1.TotalAmount())
	//cart1 := Cart{
	//	items: map["apple"]item1,
	//}
	fmt.Println(items1, "\n\n\n cart1.items is ", cart1.items["item1"].Name)
	fmt.Printf("%T", items1)
	fmt.Println("\n", customMap1)
}
