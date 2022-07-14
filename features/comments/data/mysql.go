package data

import (
	"errors"
	"fmt"
	"project/group3/features/comments"

	"gorm.io/gorm"
)

type mysqlCommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) comments.Data {
	return &mysqlCommentRepository{
		DB: db,
	}
}

func (repo *mysqlCommentRepository) InsertData(input comments.Core) (row int, err error) {
	comment := FromCore(input)
	fmt.Println("comment: ", comment)
	resultCreate := repo.DB.Create(&comment)
	fmt.Println("resultCreateerror: ", resultCreate.Error)
	if resultCreate.Error != nil {
		return 0, resultCreate.Error
	}
	if resultCreate.RowsAffected != 1 {
		return 0, errors.New("failed to insert data, your email is already registered")
	}
	return int(resultCreate.RowsAffected), nil
}

func (repo *mysqlCommentRepository) SelectCommentByIdEvent(idEvent, limitint, offsetint int) (data []comments.Core, err error) {
	dataComment := []Comment{}
	result := repo.DB.Limit(limitint).Offset(offsetint).Preload("User").Order("created_at DESC").Where("event_id = ?", idEvent).Find(&dataComment)
	if result.Error != nil {
		return []comments.Core{}, result.Error
	}
	return toCoreList(dataComment), nil
}

func (repo *mysqlCommentRepository) DeleteCommentByIdComment(idComment, idFromToken int) (row int, err error) {
	dataComment := Comment{}
	idCheck := repo.DB.First(&dataComment, idComment)
	if idCheck.Error != nil {
		return 0, idCheck.Error
	}
	if idFromToken != dataComment.UserID {
		return -1, errors.New("you don't have access")
	}
	result := repo.DB.Delete(&Comment{}, idComment)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to delete data")
	}
	return int(result.RowsAffected), nil
}
