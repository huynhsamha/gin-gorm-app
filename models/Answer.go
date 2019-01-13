package models

// Answer : Table name is `answers`
type Answer struct {
	CustomBasicModel

	QuestionID uint `json:"questionID"`

	Content string `json:"content"`

	Owner   User `gorm:"foreignkey:OwnerID"`
	OwnerID uint `json:"ownerID"`
}
