package dto

type Dropdown struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type DropdownStr struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
