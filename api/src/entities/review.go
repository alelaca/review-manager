package entities

type Review struct {
	ID      int64
	Comment string
	Rate    int
	UserID  int64
	OrderID int64
	ShopID  int64
}
