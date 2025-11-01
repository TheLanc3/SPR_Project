package parameters

type OrderCreationData struct {
	CustomerId int64
	Total      int
	Positions  []Position
}

type Position struct {
	ProductId int64
	Quantity  int
}
