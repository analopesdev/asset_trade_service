package entity

import "sync"

type Book struct {
	Order           []*Order
	Transaction     []*Transaction
	IncomingOrders  chan *Order
	ProcessedOrders chan *Order
	Wg              *sync.WaitGroup
}

func NewBook(incomingOrders chan *Order, processedOrders chan *Order, wg *sync.WaitGroup) *Book {
	return &Book{
		Order:           []*Order{},
		Transaction:     []*Transaction{},
		IncomingOrders:  incomingOrders,
		ProcessedOrders: processedOrders,
		Wg:              wg,
	}
}
