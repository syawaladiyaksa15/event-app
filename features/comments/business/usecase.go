package business

import (
	"errors"
	"project/group3/features/comments"
)

type commentUseCase struct {
	commentData comments.Data
}

func NewCommentBusiness(dataComment comments.Data) comments.Business {
	return &commentUseCase{
		commentData: dataComment,
	}
}

func (uc *commentUseCase) CreateData(input comments.Core) (row int, err error) {
	if input.EventID == 0 || input.Comment == "" {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	row, err = uc.commentData.InsertData(input)
	return row, err
}

func (uc *commentUseCase) GetCommentByIdEvent(idEvent, limitint, offsetint int) (data []comments.Core, err error) {
	data, err = uc.commentData.SelectCommentByIdEvent(idEvent, limitint, offsetint)
	return data, err
}

func (uc *commentUseCase) DeleteCommentById(idComment, idFromToken int) (row int, err error) {
	data, err := uc.commentData.DeleteCommentByIdComment(idComment, idFromToken)
	return data, err
}
