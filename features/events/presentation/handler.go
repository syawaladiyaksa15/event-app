package presentation

import (
	"fmt"
	"project/group3/features/events"
	_requestEvent "project/group3/features/events/presentation/request"
	"strconv"
	"time"

	_responseEvent "project/group3/features/events/presentation/response"
	_middlewares "project/group3/features/middlewares"
	_helper "project/group3/helper"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	EventBusiness events.Business
}

func NewEventHandler(business events.Business) *EventHandler {
	return &EventHandler{
		EventBusiness: business,
	}
}

func (h *EventHandler) CreateEvent(c echo.Context) error {
	// inisialiasi variabel dengan type struct dari request
	var newEvent _requestEvent.Event

	// binding data event
	errBind := c.Bind(&newEvent)

	validate := validator.New()
	if errValidate := validate.Struct(newEvent); errValidate != nil {
		return errValidate
	}

	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	// formfile data image event
	fileData, fileInfo, fileErr := c.Request().FormFile("image")

	// return err jika missing file
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file"))
	}

	// cek ekstension file upload
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

	// ekstrak token
	idToken, _, errToken := _middlewares.ExtractToken(c)

	// return jika errorToken
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	// inissialisasi newEvent.UserId = idToken(userid)
	newEvent.UserId = idToken
	//
	newEvent.Image = url

	dataEvent := _requestEvent.ToCore(newEvent)
	_, err := h.EventBusiness.CreateEventBusiness(dataEvent)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert data"))

	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))

}

func (h *EventHandler) DetailEvent(c echo.Context) error {
	id := c.Param("id")

	idEvent, _ := strconv.Atoi(id)

	result, err := h.EventBusiness.DetailEventBusiness(idEvent)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to detail data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseEvent.FromCore(result)))

}

func (h *EventHandler) EditEvent(c echo.Context) error {
	var editEvent _requestEvent.Event

	id := c.Param("id")

	idEvent, _ := strconv.Atoi(id)

	// validate := validator.New()
	// if errValidate := validate.Struct(editEvent); errValidate != nil {
	// 	return errValidate
	// }

	errBind := c.Bind(&editEvent)

	fmt.Println(editEvent.EventName)

	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	// formfile data image event
	fileData, fileInfo, fileErr := c.Request().FormFile("image")

	if fileErr == http.ErrMissingFile || fileErr != nil {
		// tidak ingin update image event
		// get image lama
		resImg, err := h.EventBusiness.DetailImageEventBusiness(idEvent)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to update data"))
		}

		editEvent.Image = resImg
	} else {
		// ingin update image event
		// cek ekstension file upload
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

		//
		editEvent.Image = url
	}

	idToken, _, errToken := _middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	dtEvent := _requestEvent.ToCore(editEvent)

	_, err := h.EventBusiness.UpdateEventBusiness(dtEvent, idEvent, idToken)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to update data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *EventHandler) DestroyEvent(c echo.Context) error {
	id := c.Param("id")

	idEvent, _ := strconv.Atoi(id)

	idToken, _, errToken := _middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	_, err := h.EventBusiness.DeleteEventBusiness(idEvent, idToken)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *EventHandler) JoinEvent(c echo.Context) error {
	id := c.Param("id")

	var data _requestEvent.Attendee

	errBind := c.Bind(&data)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	status := int(data.Status)

	idEvent, _ := strconv.Atoi(id)

	idToken, _, errToken := _middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	_, err := h.EventBusiness.JoinEventBusiness(idEvent, idToken, status)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to join event"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *EventHandler) AllEvent(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)

	result, err := h.EventBusiness.AllEventBusiness(limitint, offsetint)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseEvent.FromCoreList(result)))

}

func (h *EventHandler) MyEvent(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)

	idToken, _, errToken := _middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	result, err := h.EventBusiness.MyEventBusiness(limitint, offsetint, idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseEvent.FromCoreList(result)))

}

func (h *EventHandler) AttendeeEvent(c echo.Context) error {
	id := c.Param("id")

	idEvent, _ := strconv.Atoi(id)

	result, err := h.EventBusiness.AttendeeEventBusiness(idEvent)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", result))

}
