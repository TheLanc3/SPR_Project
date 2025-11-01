package enums

type OrderType int

const (
	All OrderType = iota
	OnlyUnfinished
	OnlyDelivered
)
