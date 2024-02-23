package dto

import "time"

type Base struct {
	IsActived bool      `json:"is_actived"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
