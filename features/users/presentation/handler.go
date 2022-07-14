package presentation

import (
	"fmt"
	"net/http"
	"time"

	_middleware "project/group3/features/middlewares"
	"project/group3/features/users"
	_requestUser "project/group3/features/users/presentation/request"
	_responseUser "project/group3/features/users/presentation/response"

	_helper "project/group3/helper"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(userBusiness users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: userBusiness,
	}
}

func (h *UserHandler) PostUser(c echo.Context) error {
	userReq := _requestUser.User{}
	err := c.Bind(&userReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	dataUser := _requestUser.ToCore(userReq)
	row, errCreate := h.userBusiness.CreateData(dataUser)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
	}
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("your email is already registered"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *UserHandler) LoginAuth(c echo.Context) error {
	authData := users.AuthRequestData{}
	c.Bind(&authData)
	token, name, avatarUrl, e := h.userBusiness.LoginUser(authData)
	if e != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("email or password incorrect"))
	}

	data := map[string]interface{}{
		"token":     token,
		"name":      name,
		"avatarUrl": avatarUrl,
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("login success", data))
}

func (h *UserHandler) PutUser(c echo.Context) error {
	idFromToken, _, _ := _middleware.ExtractToken(c)
	userReq := _requestUser.User{}
	err := c.Bind(&userReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("avatar_url")
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file"))
	}

	extension, err_check_extension := _helper.CheckFileExtension(fileInfo.Filename)
	if err_check_extension != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file extension error"))
	}

	// check file size
	err_check_size := _helper.CheckFileSize(fileInfo.Size)
	if err_check_size != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file size error"))
	}

	// memberikan nama file
	fileName := time.Now().Format("2006-01-02 15:04:05") + "." + extension

	url, errUploadImg := _helper.UploadImageToS3(fileName, fileData)

	if errUploadImg != nil {
		fmt.Println(errUploadImg)
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to upload file"))
	}
	userReq.AvatarUrl = url

	dataUser := _requestUser.ToCore(userReq)
	row, errUpd := h.userBusiness.UpdateData(dataUser, idFromToken)
	if errUpd != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to update data users"))
	}
	if row == 0 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to update data users"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *UserHandler) GetByMe(c echo.Context) error {
	idFromToken, _, _ := _middleware.ExtractToken(c)
	result, errGet := h.userBusiness.GetUserByMe(idFromToken)
	if errGet != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get data user"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseUser.FromCore(result)))
}

func (h *UserHandler) DeleteByID(c echo.Context) error {
	idFromToken, _, _ := _middleware.ExtractToken(c)

	row, errDel := h.userBusiness.DeleteDataById(idFromToken)
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete data user"))
	}
	if row != 1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to delete data user"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}
