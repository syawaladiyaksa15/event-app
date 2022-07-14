package response

import (
	"project/group3/features/comments"
)

type Comment struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func FromCore(data comments.Core) Comment {
	return Comment{
		ID:      data.ID,
		Name:    data.UserName,
		Comment: data.Comment,
	}
}

func FromCoreList(data []comments.Core) []Comment {
	result := []Comment{}
	for k, _ := range data {
		result = append(result, FromCore(data[k]))
	}
	return result
}
