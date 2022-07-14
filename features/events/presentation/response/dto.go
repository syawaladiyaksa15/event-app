package response

import (
	"project/group3/features/events"
	"time"
)

type Event struct {
	ID          int       `json:"id" form:"id"`
	Image       string    `json:"image" form:"image"`
	EventName   string    `json:"event_name" form:"event_name"`
	Category    string    `json:"category" form:"category"`
	Link        string    `json:"link" form:"link"`
	Lat         string    `json:"lat" form:"lat"`
	Long        string    `json:"long" form:"long"`
	Quota       uint      `json:"quota" form:"quota"`
	Date        string    `json:"date" form:"date"`
	Time        string    `json:"time" form:"time"`
	Description string    `json:"description" form:"description"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	User        User      `json:"user" form:"user"`
}

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

func FromCore(data events.Core) Event {
	return Event{
		ID:          data.ID,
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
		User: User{
			ID:        data.User.ID,
			Name:      data.User.Name,
			AvatarUrl: data.User.AvatarUrl,
		},
	}
}

func FromCoreList(data []events.Core) []Event {
	result := []Event{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
