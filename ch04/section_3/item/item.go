package item

type Item struct {
	name     string
	price    float64
	quantity int
}

func (t *Item) Name() string {
	return t.name
}

func (t *Item) SetName(n string) {
	if len(n) != 0 {
		t.name = n
	}
}

func (t *Item) Price() float64 {
	return t.price
}

func (t *Item) SetPrice(p float64) {
	if p > 0 {
		t.price = p
	}
}

func (t *Item) Quantity() int {
	return t.quantity
}
func (t *Item) SetQuantity(q int) {
	if q > 0 {
		t.quantity = q
	}
}

func New(name string, price float64, quantity int) *Item {
	if price <= 0 || quantity <= 0 || len(name) == 0 {
		return nil
	}
	return &Item{name, price, quantity}
}
