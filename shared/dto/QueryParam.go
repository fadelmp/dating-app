package dto

import "time"

type QueryParam struct {
	Search string    `query:"search"`
	Page   uint      `query:"page"`
	Limit  uint      `query:"limit"`
	Sort   string    `query:"sort"`
	Order  string    `query:"order"`
	From   time.Time `query:"from"`
	To     time.Time `query:"to"`
}
