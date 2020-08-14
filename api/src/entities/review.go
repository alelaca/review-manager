package entities

import (
	"time"
)

type Review struct {
	ID              *int64     `json:"id"`
	Comment         *string    `json:"comment"`
	Rate            *int       `json:"rate"`
	UserID          *int64     `json:"user_id"`
	OrderID         *int64     `json:"order_id"`
	ShopID          *int64     `json:"shop_id"`
	DateCreated     *time.Time `json:"date_created"`
	DateLastUpdated *time.Time `json:"date_last_updated"`
}
