package request

import (
	"project/group3/features/events"
)

type Event struct {
	Image       string `json:"image" form:"image"`
	EventName   string `json:"event_name" form:"event_name" validate:"required,min=5"`
	Category    string `json:"category" form:"category" validate:"required"`
	Link        string `json:"link" form:"link"`
	Lat         string `json:"lat" form:"lat"`
	Long        string `json:"long" form:"long"`
	Quota       uint   `json:"quota" form:"quota" validate:"required,number"`
	Date        string `json:"date" form:"date" validate:"required"`
	Time        string `json:"time" form:"time" validate:"required"`
	Description string `json:"description" form:"description"`
	UserId      int    `json:"user_id" form:"user_id"`
}

type Attendee struct {
	Status uint `json:"status" form:"status"`
}

func ToCore(req Event) events.Core {
	return events.Core{
		Image:       req.Image,
		EventName:   req.EventName,
		Category:    req.Category,
		Link:        req.Link,
		Lat:         req.Lat,
		Long:        req.Long,
		Quota:       req.Quota,
		Date:        req.Date,
		Time:        req.Time,
		Description: req.Description,
		User: events.User{
			ID: req.UserId,
		},
	}
}
