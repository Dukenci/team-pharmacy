package models

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type CategoryCreateRequest struct {
	Name string `json:"name"`
}

type CategoryUpdateRequest struct {
	Name *string `json:"name"`
}
