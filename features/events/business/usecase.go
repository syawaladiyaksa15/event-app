package business

import (
	"errors"
	"project/group3/features/events"
)

type eventUseCase struct {
	eventData events.Data
}

func NewEventBusiness(evtData events.Data) events.Business {
	return &eventUseCase{
		eventData: evtData,
	}
}

func (uc *eventUseCase) CreateEventBusiness(newData events.Core) (response int, err error) {
	response, err = uc.eventData.InsertData(newData)

	return response, err
}

func (uc *eventUseCase) DetailEventBusiness(id int) (response events.Core, err error) {
	response, err = uc.eventData.DetailEventData(id)

	return response, err
}

func (uc *eventUseCase) UpdateEventBusiness(editData events.Core, id int, idUser int) (response int, err error) {
	if editData.EventName == "" || editData.Category == "" || editData.Quota == 0 || editData.Date == "" || editData.Description == "" || editData.Time == "" {
		return 0, errors.New("all input data must be filled")
	}

	response, err = uc.eventData.UpdateEventData(editData, id, idUser)

	return response, err
}

func (uc *eventUseCase) DetailImageEventBusiness(id int) (response string, err error) {
	response, err = uc.eventData.DetailImageEventData(id)

	return response, err
}

func (uc *eventUseCase) DeleteEventBusiness(id int, idUser int) (row int, err error) {
	row, err = uc.eventData.DeleteEventData(id, idUser)

	return row, err
}

func (uc *eventUseCase) JoinEventBusiness(id, idUser, status int) (row int, err error) {
	row, err = uc.eventData.JoinEventData(id, idUser, status)

	return row, err
}

func (uc *eventUseCase) AllEventBusiness(limit, offset int) (response []events.Core, err error) {
	response, err = uc.eventData.AllEventData(limit, offset)

	return response, err
}

func (uc *eventUseCase) MyEventBusiness(limit, offset, idUser int) (response []events.Core, err error) {
	response, err = uc.eventData.MyEventData(limit, offset, idUser)

	return response, err
}

func (uc *eventUseCase) AttendeeEventBusiness(id int) (response []events.User, err error) {
	response, err = uc.eventData.AttendeeEventData(id)

	return response, err
}
