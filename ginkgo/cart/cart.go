package cart

type Cart struct {
	items map[string]Item
}
type Item struct {
	ID    string
	Name  string
	Price float64
	Qty   int
}
