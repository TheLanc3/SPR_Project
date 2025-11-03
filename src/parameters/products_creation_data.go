package parameters

type Product struct {
	Name        string
	Description string
	SupplierId  int64
	Price       int
	Quantity    int
}

type ProductUpdate struct {
	Id          int64
	Description string
	Price       int
}
