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

type orderQueue []*Order

func (oq orderQueue) Add(order *Order) orderQueue {
	return append(oq, order)
}

func (oq *orderQueue) GetNextOrder() *Order {
	if len(*oq) == 0 {
		return nil
	}

	order := (*oq)[0]
	*oq = (*oq)[1:]

	return order
}

func (b *Book) Trade() {
	buyOrders := make(map[string]*orderQueue)
	sellOrders := make(map[string]*orderQueue)

	for order := range b.IncomingOrders {
		asset := order.Asset.ID

		if buyOrders[asset] == nil {
			buyOrders[asset] = &orderQueue{}
		}

		if sellOrders[asset] == nil {
			sellOrders[asset] = &orderQueue{}
		}

		if order.OrderType == "BUY" {
			b.tryMatch(order, sellOrders[asset], buyOrders[asset])
		} else {
			b.tryMatch(order, buyOrders[asset], sellOrders[asset])
		}

	}
}

func (b *Book) tryMatch(newOrder *Order, availableOrders, pendingOrders *orderQueue) {
	for {
		potentialMatches := availableOrders.GetNextOrder()

		if potentialMatches == nil {
			break
		}

		if !b.pricesMatch(newOrder, potentialMatches) {
			availableOrders.Add(potentialMatches)
			break
		}

		if potentialMatches.PendingShares > 0 {
			matchedTransaction := b.createTransaction(newOrder, potentialMatches)
			b.processTransaction(matchedTransaction)

			if potentialMatches.PendingShares > 0 {
				availableOrders.Add(potentialMatches)
			}

			if newOrder.PendingShares == 0 {
				break
			}
		}
	}

	if newOrder.PendingShares > 0 {
		availableOrders.Add(newOrder)
	}
}

func (b *Book) pricesMatch(order, matchOrder *Order) bool {
	if order.OrderType == "BUY" {
		return matchOrder.Price <= order.Price
	}

	return order.Price <= matchOrder.Price
}

func (b *Book) createTransaction(incomingOrder, matchedOrder *Order) *Transaction {
	var buyOrder, sellOrder *Order

	if incomingOrder.OrderType == "BUY" {
		buyOrder, sellOrder = incomingOrder, matchedOrder
	} else {
		buyOrder, sellOrder = matchedOrder, incomingOrder
	}

	shares := incomingOrder.PendingShares

	if matchedOrder.PendingShares < shares {
		shares = matchedOrder.PendingShares
	}

	return NewTransation(sellOrder, buyOrder, shares, matchedOrder.Price)
}

func (b *Book) recordTransaction(transaction *Transaction) {
	b.Transaction = append(b.Transaction, transaction)
	transaction.BuyingOrder.Transactions = append(transaction.BuyingOrder.Transactions, transaction)
	transaction.SellingOrder.Transactions = append(transaction.SellingOrder.Transactions, transaction)
}

func (b *Book) processTransaction(transaction *Transaction) {
	defer b.Wg.Done()

	transaction.Process()
	b.recordTransaction(transaction)
	b.ProcessedOrders <- transaction.BuyingOrder
	b.ProcessedOrders <- transaction.SellingOrder
}
