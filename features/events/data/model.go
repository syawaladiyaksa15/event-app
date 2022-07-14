package data

import (
	"project/group3/features/events"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	// gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement"`
	UserID      uint `json:"user_id" form:"user_id"`
	User        User
	Image       string `json:"image" form:"image" gorm:"not null; type:varchar(255)"`
	EventName   string `json:"event_name" form:"event_name" gorm:"not null; type:varchar(255); unique"`
	Category    string `json:"category" form:"category" gorm:"not null; type:varchar(100)"`
	Link        string `json:"link" form:"link" gorm:"type:text"`
	Lat         string `json:"lat" form:"lat" gorm:"type:varchar(255)"`
	Long        string `json:"long" form:"long" gorm:"type:varchar(255)"`
	Quota       uint   `json:"quota" form:"quota" gorm:"not null; type:integer"`
	Date        string `json:"date" form:"date" gorm:"not null; type:date"`
	Time        string `json:"time" form:"time" gorm:"not null; type:varchar(10)"`
	Description string `json:"description" form:"description" gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Attendee    []Attendee     `gorm:"foreignKey:EventID;references:ID;constraint:OnDelete:CASCADE"`
}

type Attendee struct {
	// gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint `json:"user_id" form:"user_id"`
	User      User
	EventID   uint `json:"event_id" form:"event_id"`
	Event     Event
	Status    uint `json:"status" form:"status" gorm:"type:integer; default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	gorm.Model
	Name      string     `json:"name" form:"name"`
	AvatarUrl string     `json:"avatar_url" form:"avatar_url"`
	Events    []Event    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Attendees []Attendee `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

func (data *Event) toCore() events.Core {
	return events.Core{
		ID: int(data.ID),
		User: events.User{
			ID:        int(data.User.ID),
			Name:      data.User.Name,
			AvatarUrl: data.User.AvatarUrl,
		},
		Image:       data.Image,
		EventName:   data.EventName,
		Category:    data.Category,
		Link:        data.Link,
		Lat:         data.Lat,
		Long:        data.Long,
		Quota:       data.Quota,
		Date:        data.Date,
		Time:        data.Time,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func (data *User) toUser() events.User {
	return events.User{
		ID:        int(data.ID),
		Name:      data.Name,
		AvatarUrl: data.AvatarUrl,
	}
}

func toCoreList(data []Event) []events.Core {
	result := []events.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func toCoreListUser(data []User) []events.User {
	result := []events.User{}
	for key := range data {
		result = append(result, data[key].toUser())
	}
	return result
}

func fromCore(core events.Core) Event {
	return Event{
		Image:       core.Image,
		EventName:   core.EventName,
		Category:    core.Category,
		Link:        core.Link,
		Lat:         core.Lat,
		Long:        core.Long,
		Quota:       core.Quota,
		Date:        core.Date,
		Time:        core.Time,
		Description: core.Description,
		UserID:      uint(core.User.ID),
	}
}
