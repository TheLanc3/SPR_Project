package parameters

type OrderCreationData struct {
	CustomerId int64
	Total      int
	Positions  []Position
}

func NewOrderCreationData(customerId int64,
	positions []Position) OrderCreationData {
	var total int

	for _, position := range positions {
		total += position.Quantity
	}

	return OrderCreationData{customerId, total, positions}
}

type Position struct {
	ProductId int64
	Price     int
	Quantity  int
}
