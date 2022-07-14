package factory

import (
	_userBusiness "project/group3/features/users/business"
	_userData "project/group3/features/users/data"
	_userPresentation "project/group3/features/users/presentation"

	_eventBusiness "project/group3/features/events/business"
	_eventData "project/group3/features/events/data"
	_eventPresentation "project/group3/features/events/presentation"

	_commentBusiness "project/group3/features/comments/business"
	_commentData "project/group3/features/comments/data"
	_commentPresentation "project/group3/features/comments/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter    *_userPresentation.UserHandler
	EventPresenter   *_eventPresentation.EventHandler
	CommentPresenter *_commentPresentation.CommentHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	// dbConn := config.InitDB()
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	eventData := _eventData.NewEventRepository(dbConn)
	eventBusiness := _eventBusiness.NewEventBusiness(eventData)
	eventPresentation := _eventPresentation.NewEventHandler(eventBusiness)

	commentData := _commentData.NewCommentRepository(dbConn)
	commentBusiness := _commentBusiness.NewCommentBusiness(commentData)
	commentPresentation := _commentPresentation.NewCommentHandler(commentBusiness)

	return Presenter{
		UserPresenter:    userPresentation,
		EventPresenter:   eventPresentation,
		CommentPresenter: commentPresentation,
	}
}
