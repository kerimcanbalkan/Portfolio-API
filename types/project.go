package types

type ProjectCreateRequest = struct {
	Title       string `json:"title" gorm:"not null;unique"`
	Image       string `json:"image" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
}
