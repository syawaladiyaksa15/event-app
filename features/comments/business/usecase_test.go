package business

import (
	"fmt"
	"project/group3/features/comments"
	"testing"

	"github.com/stretchr/testify/assert"
)

// success

type mockCommentData struct{}

func (mock mockCommentData) InsertData(input comments.Core) (row int, err error) {
	return 1, nil
}

func (mock mockCommentData) SelectCommentByIdEvent(idEvent, limitint, offsetint int) (data []comments.Core, err error) {
	return []comments.Core{
		{
			ID:       19,
			UserName: "dwi",
			Comment:  "keduax"},
		{
			ID:       18,
			UserName: "dwi1",
			Comment:  "pertamax",
		},
	}, nil
}

func (mock mockCommentData) DeleteCommentByIdComment(idComment, idFromToken int) (row int, err error) {
	return 1, nil
}

// failed

type mockCommentDataFailed struct{}

func (mock mockCommentDataFailed) InsertData(input comments.Core) (row int, err error) {
	return 0, fmt.Errorf("failed to insert data ")
}

func (mock mockCommentDataFailed) SelectCommentByIdEvent(idEvent, limitint, offsetint int) (data []comments.Core, err error) {
	return []comments.Core{}, fmt.Errorf("failed to get data comment")
}

func (mock mockCommentDataFailed) DeleteCommentByIdComment(idComment, idFromToken int) (row int, err error) {
	return 0, fmt.Errorf("failed to delete data comment")
}

func TestInsertData(t *testing.T) {
	t.Run("Test InsertData Success", func(t *testing.T) {
		input := comments.Core{
			EventID: 1,
			Comment: "pertamax",
		}
		commentBusiness := NewCommentBusiness(mockCommentData{})
		result, err := commentBusiness.CreateData(input)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Test InsertData Failed", func(t *testing.T) {
		input := comments.Core{
			EventID: 1,
			Comment: "pertamax",
		}
		commentBusiness := NewCommentBusiness(mockCommentDataFailed{})
		result, err := commentBusiness.CreateData(input)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestSelectCommentByIdEvent(t *testing.T) {
	t.Run("Test SelectCommentByIdEvent Success", func(t *testing.T) {
		idEvent := 1
		limitint := 0
		offsetint := 0
		commentBusiness := NewCommentBusiness(mockCommentData{})
		result, err := commentBusiness.GetCommentByIdEvent(idEvent, limitint, offsetint)
		assert.Nil(t, err)
		assert.Equal(t, []comments.Core{
			{
				ID:       19,
				UserName: "dwi",
				Comment:  "keduax"},
			{
				ID:       18,
				UserName: "dwi1",
				Comment:  "pertamax",
			},
		}, result)
	})
	t.Run("Test SelectCommentByIdEvent Failed", func(t *testing.T) {
		idEvent := 1
		limitint := 0
		offsetint := 0
		commentBusiness := NewCommentBusiness(mockCommentDataFailed{})
		result, err := commentBusiness.GetCommentByIdEvent(idEvent, limitint, offsetint)
		assert.NotNil(t, err)
		assert.Equal(t, []comments.Core{}, result)
	})

}

func TestDeleteCommentByIdComment(t *testing.T) {
	t.Run("Test DeleteCommentById Success", func(t *testing.T) {
		idComment := 19
		idFromToken := 1
		commentBusiness := NewCommentBusiness(mockCommentData{})
		result, err := commentBusiness.DeleteCommentById(idComment, idFromToken)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Test DeleteCommentById Failed", func(t *testing.T) {
		idComment := 19
		idFromToken := 1
		commentBusiness := NewCommentBusiness(mockCommentDataFailed{})
		result, err := commentBusiness.DeleteCommentById(idComment, idFromToken)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}
