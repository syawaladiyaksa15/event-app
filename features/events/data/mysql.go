package data

import (
	"fmt"
	"project/group3/features/events"

	"gorm.io/gorm"
)

type mysqlEventRepository struct {
	db *gorm.DB
}

func NewEventRepository(conn *gorm.DB) events.Data {
	return &mysqlEventRepository{
		db: conn,
	}
}

func (repo *mysqlEventRepository) InsertData(input events.Core) (response int, err error) {
	event := fromCore(input)
	result := repo.db.Create(&event)

	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}

	return int(result.RowsAffected), err
}

func (repo *mysqlEventRepository) DetailEventData(id int) (response events.Core, err error) {
	var dataEvent Event

	result := repo.db.Preload("User").First(&dataEvent, "id = ?", id)

	if result.RowsAffected != 1 {
		return events.Core{}, fmt.Errorf("event not found")
	}

	if result.Error != nil {
		return events.Core{}, result.Error
	}

	return dataEvent.toCore(), nil
}

func (repo *mysqlEventRepository) UpdateEventData(editData events.Core, id int, idUser int) (response int, err error) {

	event := fromCore(editData)
	event_ := fromCore(editData)

	searchEvent := repo.db.First(&event_, "id = ?", id)

	if searchEvent.RowsAffected != 1 {
		return 0, fmt.Errorf("failed update event")
	}

	if event_.UserID != uint(idUser) {
		return 0, fmt.Errorf("failed update event")
	}

	result := repo.db.Model(Event{}).Where("id = ?", id).Updates(&event)

	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("event not found")
	}

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (repo *mysqlEventRepository) DetailImageEventData(id int) (response string, err error) {
	var dataEvent Event

	result := repo.db.Preload("User").First(&dataEvent, "id = ?", id)

	if result.RowsAffected != 1 {
		return "", fmt.Errorf("event not found")
	}

	if result.Error != nil {
		return "", result.Error
	}

	return dataEvent.Image, nil
}

func (repo *mysqlEventRepository) DeleteEventData(id int, idUser int) (row int, err error) {
	var dataEvent Event

	searchProduct := repo.db.Find(&dataEvent, id)

	if searchProduct.RowsAffected != 1 {
		return 0, fmt.Errorf("failed delete event")
	}

	if searchProduct.Error != nil {
		return 0, searchProduct.Error
	}

	if dataEvent.UserID != uint(idUser) {
		return 0, fmt.Errorf("failed delete event")
	}

	result := repo.db.Delete(&dataEvent, id)

	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("event not found")
	}

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (repo *mysqlEventRepository) JoinEventData(id, idUser, status int) (row int, err error) {
	var dataEvent Event

	// pengecekan pemilik event
	searchEvent := repo.db.First(&dataEvent, id)

	if searchEvent.Error != nil || searchEvent.RowsAffected != 1 {
		return 0, fmt.Errorf("event not found")
	}

	if dataEvent.UserID == uint(idUser) {
		return 0, fmt.Errorf("join event failed")
	}

	if status == 1 {
		// pengecekan jumlah member event
		var countAttendee int64
		var countAttendee_ int64

		rsCount := repo.db.Table("attendees").Where("event_id = ? AND status = ?", id, status).Count(&countAttendee)

		if rsCount.Error != nil {
			return 0, fmt.Errorf("join event failed")
		}

		if dataEvent.Quota <= uint(countAttendee) {
			return 0, fmt.Errorf("join event failed")
		}

		// pengecekan apakah user sudah join sebelumnya
		rsCount_ := repo.db.Table("attendees").Where("event_id = ? AND user_id = ?", id, idUser).Count(&countAttendee_)

		if rsCount_.Error != nil {
			return 0, fmt.Errorf("join event failed")
		}

		if countAttendee_ >= 1 {
			// return 0, fmt.Errorf("join event failed")
			// update
			rsUpd := repo.db.Model(Attendee{}).Where("event_id = ? AND user_id = ?", id, idUser).Update("status", status)

			if rsUpd.RowsAffected != 1 {
				return 0, fmt.Errorf("join event failed")
			}

			if rsUpd.Error != nil {
				return 0, fmt.Errorf("join event failed")
			}

			return int(rsUpd.RowsAffected), nil

		}

		var dataAttendee Attendee

		dataAttendee.UserID = uint(idUser)
		dataAttendee.EventID = uint(id)
		dataAttendee.Status = uint(status)

		result := repo.db.Create(&dataAttendee)

		if result.Error != nil || result.RowsAffected != 1 {
			return 0, fmt.Errorf("join event failed")
		}

		return int(result.RowsAffected), nil

	} else {
		// langsung cancel join event
		// pengecekkan userid dan eventid
		var checkData int64

		rsCheck := repo.db.Table("attendees").Where("event_id = ? AND user_id = ?", id, idUser).Count(&checkData)

		if rsCheck.Error != nil {
			return 0, fmt.Errorf("cancel join event failed")
		}

		if checkData < 1 {
			return 0, fmt.Errorf("cancel join event failed")
		}

		// update
		result := repo.db.Model(Attendee{}).Where("event_id = ? AND user_id = ?", id, idUser).Update("status", status)

		if result.RowsAffected != 1 {
			return 0, fmt.Errorf("cancel join event failed")
		}

		if result.Error != nil {
			return 0, fmt.Errorf("cancel join event failed")
		}

		return int(result.RowsAffected), nil
	}
}

func (repo *mysqlEventRepository) AllEventData(limit, offset int) (response []events.Core, err error) {
	var dataEvents []Event

	result := repo.db.Preload("User").Order("id desc").Limit(limit).Offset(offset).Find(&dataEvents)

	if result.Error != nil {
		return []events.Core{}, result.Error
	}

	return toCoreList(dataEvents), nil
}

func (repo *mysqlEventRepository) MyEventData(limit, offset, idUser int) (response []events.Core, err error) {
	var dataEvents []Event

	result := repo.db.Preload("User").Order("id desc").Limit(limit).Offset(offset).Find(&dataEvents, "user_id = ?", idUser)

	if result.Error != nil {
		return []events.Core{}, result.Error
	}

	return toCoreList(dataEvents), nil
}

func (repo *mysqlEventRepository) AttendeeEventData(id int) (response []events.User, err error) {
	var data []Attendee
	var dataUser []events.User

	var countCheck int64
	checkEvent := repo.db.Table("events").Where("id = ?", id).Count(&countCheck)

	if checkEvent.Error != nil {
		return []events.User{}, fmt.Errorf("not found event")
	}

	if countCheck < 1 {
		return []events.User{}, fmt.Errorf("not found event")
	}

	result := repo.db.Preload("User").Table("attendees").Where("event_id = ? AND status = ?", id, 1).Find(&data)

	if result.Error != nil {
		return []events.User{}, fmt.Errorf("failed show data")
	}

	for k := range data {
		dataUser = append(dataUser, data[k].User.toUser())
	}

	return dataUser, nil
}

// func (repo *mysqlEventRepository) AttendeeEventData(id int) (response []map[string]interface{}, err error) {
// 	var data []map[string]interface{}

// 	result := repo.db.Table("attendees").Where("event_id = ? AND status = ?", id, 1).Find(&data)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	// var formatData []map[string]interface{}
// 	var formatData []User

// 	for key := range data {
// 		var dataUser events.User

// 		rs := repo.db.Table("users").Where("id = ?", data[key]["user_id"]).First(&dataUser)

// 		if rs.Error != nil {
// 			return nil, rs.Error
// 		}

// 		formatData[key].ID = uint(data[key]["user_id"].(int64))
// 		formatData[key].Name = dataUser.Name
// 		formatData[key].AvatarUrl = dataUser.AvatarUrl
// 	}

// 	// fmt.Println(formatData)
// 	fmt.Println(data)

// 	return data, nil
// }
