package presentation

import (
	"net/http"
	"project/group3/features/comments"
	_requestComment "project/group3/features/comments/presentation/request"
	_responseComment "project/group3/features/comments/presentation/response"
	_middleware "project/group3/features/middlewares"
	_helper "project/group3/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentBusiness comments.Business
}

func NewCommentHandler(commentBusiness comments.Business) *CommentHandler {
	return &CommentHandler{
		commentBusiness: commentBusiness,
	}
}

func (h *CommentHandler) PostComment(c echo.Context) error {
	id := c.Param("id")
	idEvent, _ := strconv.Atoi(id)
	commentReq := _requestComment.Comment{}
	err := c.Bind(&commentReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}
	idFromToken, _, _ := _middleware.ExtractToken(c)
	dataComment := _requestComment.ToCore(commentReq)
	dataComment.UserID = idFromToken
	dataComment.EventID = idEvent
	row, errCreate := h.commentBusiness.CreateData(dataComment)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
	}

	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to add comment"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *CommentHandler) GetComment(c echo.Context) error {
	id := c.Param("id")
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	idEvent, _ := strconv.Atoi(id)
	result, errGet := h.commentBusiness.GetCommentByIdEvent(idEvent, limitint, offsetint)
	if errGet != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get data user"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseComment.FromCoreList(result)))
}

func (h *CommentHandler) DeleteComment(c echo.Context) error {
	id := c.Param("id")
	idComment, _ := strconv.Atoi(id)
	idFromToken, _, _ := _middleware.ExtractToken(c)
	row, errDel := h.commentBusiness.DeleteCommentById(idComment, idFromToken)
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete data user"))
	}
	if row != 1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to delete data user"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}
