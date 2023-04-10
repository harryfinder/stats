package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и тд)
type Money int64

// Category представляет собой категорию, в которой был совершён платёж (авто, аптекы,рестораны и тд)
type Category string

// Payment представляет информацию о платеже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}

// PaymentSource представяет из себя источник оплаты
type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}

// Card представдяет информацию о платежной карте
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}

// Currency представляет собой денежную единицу
type Currency string

// коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

// PAN представляет номер карты
type PAN string
