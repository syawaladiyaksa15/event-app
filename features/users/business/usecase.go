package business

import (
	"errors"
	"fmt"
	"project/group3/features/users"

	_bcrypt "golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userData users.Data
}

func NewUserBusiness(dataUser users.Data) users.Business {
	return &userUseCase{
		userData: dataUser,
	}
}

func (uc *userUseCase) CreateData(input users.Core) (row int, err error) {
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	row, err = uc.userData.InsertData(input)
	return row, err
}

func (uc *userUseCase) LoginUser(authData users.AuthRequestData) (token, name, avatarUrl string, err error) {
	token, name, avatarUrl, err = uc.userData.LoginUserDB(authData)
	return token, name, avatarUrl, err
}

func (uc *userUseCase) UpdateData(input users.Core, idFromToken int) (row int, err error) {
	userReq := map[string]interface{}{}
	if input.Name != "" {
		userReq["name"] = input.Name
	}
	if input.Email != "" {
		userReq["email"] = input.Email
	}
	if input.Password != "" {
		passwordHashed, errorHash := _bcrypt.GenerateFromPassword([]byte(input.Password), 10)
		if errorHash != nil {
			fmt.Println("Error hash", errorHash.Error())
		}
		userReq["password"] = string(passwordHashed)
	}
	if input.AvatarUrl != "" {
		userReq["avatar_url"] = input.AvatarUrl
	}
	row, err = uc.userData.UpdateDataDB(userReq, idFromToken)
	return row, err
}

func (uc *userUseCase) GetUserByMe(idFromToken int) (data users.Core, err error) {
	data, err = uc.userData.SelectDataByMe(idFromToken)
	return data, err
}

func (uc *userUseCase) DeleteDataById(idFromToken int) (row int, err error) {
	row, err = uc.userData.DeleteDataByIdDB(idFromToken)
	return row, err
}
