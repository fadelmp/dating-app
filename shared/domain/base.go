package domain

import "time"

type Base struct {
	IsActived bool      `json:"is_actived"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

func Create(name string) Base {

	return Base{
		IsActived: true,
		IsDeleted: false,
		CreatedBy: name,
		UpdatedBy: name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func Update(name string) Base {

	return Base{
		UpdatedBy: name,
		UpdatedAt: time.Now(),
	}
}

func Delete(name string) Base {

	return Base{
		IsActived: false,
		IsDeleted: true,
		UpdatedBy: name,
		UpdatedAt: time.Now(),
	}
}
