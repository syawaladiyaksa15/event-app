package request

import "project/group3/features/comments"

type Comment struct {
	EventID int    `json:"event_id" form:"event_id"`
	Comment string `json:"comment" form:"comment"`
}

func ToCore(req Comment) comments.Core {
	return comments.Core{
		EventID: req.EventID,
		Comment: req.Comment,
	}
}
