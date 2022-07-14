package users

import "time"

type Core struct {
	ID        int
	Name      string
	Email     string
	Password  string
	AvatarUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	CreateData(input Core) (row int, err error)
	LoginUser(authData AuthRequestData) (token, name, avatarUrl string, err error)
	UpdateData(input Core, idUser int) (row int, err error)
	GetUserByMe(idFromToken int) (data Core, err error)
	DeleteDataById(idFromToken int) (row int, err error)
}

type Data interface {
	InsertData(input Core) (row int, err error)
	LoginUserDB(authData AuthRequestData) (token, name, avatarUrl string, err error)
	UpdateDataDB(data map[string]interface{}, idUser int) (row int, err error)
	SelectDataByMe(idFromToken int) (data Core, err error)
	DeleteDataByIdDB(idFromToken int) (row int, err error)
}
