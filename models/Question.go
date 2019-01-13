package models

// Question : Table name is `questions`
type Question struct {
	CustomBasicModel

	Title   string `json:"title"`
	Content string `json:"content"`

	Owner   User `gorm:"foreignkey:OwnerID"`
	OwnerID uint `json:"ownerID"`

	Votes uint `gorm:"default:0" json:"votes"`

	AcceptedAnswerID uint     `json:"acceptedAnswerID"`
	Answers          []Answer `gorm:"foreignkey:QuestionID;association_foreignkey:ID"`
}
