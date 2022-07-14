package comments

import "time"

type Core struct {
	ID        int
	EventID   int
	UserID    int
	UserName  string
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  time.Time
	User      User
}

type User struct {
	ID   int
	Name string
}

type Business interface {
	CreateData(input Core) (row int, err error)
	GetCommentByIdEvent(idEvent, limitint, offsetint int) (data []Core, err error)
	DeleteCommentById(idComment, idFromToken int) (row int, err error)
}

type Data interface {
	InsertData(input Core) (row int, err error)
	SelectCommentByIdEvent(idEvent, limitint, offsetint int) (data []Core, err error)
	DeleteCommentByIdComment(idComment, idFromToken int) (row int, err error)
}
