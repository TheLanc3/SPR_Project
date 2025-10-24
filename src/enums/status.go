package enums

type Status int

const (
	Ordered Status = iota
	Shipped
	Delivered
	Completed
	Canceled
)
