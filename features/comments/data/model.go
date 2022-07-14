package data

import (
	"project/group3/features/comments"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	EventID int
	UserID  int
	Comment string
	User    User  `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Events  Event `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type User struct {
	gorm.Model
	Name      string
	AvatarUrl string
	Comment   []Comment
}

type Event struct {
	gorm.Model
}

func (data *Comment) toCore() comments.Core {
	return comments.Core{
		ID:       int(data.ID),
		UserID:   int(data.UserID),
		UserName: data.User.Name,
		User: comments.User{
			ID:   int(data.User.ID),
			Name: data.User.Name,
		},
		EventID:   int(data.EventID),
		Comment:   data.Comment,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func toCoreList(data []Comment) []comments.Core {
	result := []comments.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func FromCore(core comments.Core) Comment {
	return Comment{
		EventID: int(core.EventID),
		UserID:  int(core.UserID),
		Comment: core.Comment,
	}
}
