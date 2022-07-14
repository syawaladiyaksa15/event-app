package migration

import (
	_mComments "project/group3/features/comments/data"
	_mEvents "project/group3/features/events/data"
	_mUsers "project/group3/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})
	db.AutoMigrate(&_mEvents.Event{})
	db.AutoMigrate(&_mEvents.Attendee{})
	db.AutoMigrate(&_mComments.Comment{})
}
