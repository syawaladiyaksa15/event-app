package data

import (
	"project/group3/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	AvatarUrl string
}

// DTO

func (data *User) toCore() users.Core {
	return users.Core{
		ID:        int(data.ID),
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		AvatarUrl: data.AvatarUrl,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func toCoreList(data []User) []users.Core {
	result := []users.Core{}
	for k := range data {
		result = append(result, data[k].toCore())
	}
	return result
}

func FromCore(core users.Core) User {
	return User{
		Name:      core.Name,
		Email:     core.Email,
		Password:  core.Password,
		AvatarUrl: core.AvatarUrl,
	}
}
