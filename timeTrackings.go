package ebrest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type TimeTrackingsResource struct {
	client   *Client
	endpoint string
}

func (TimeTracking *TimeTrackingsResource) GetTimeTrackings(params map[string]string) (PaginatedResponse, error) {
	var payload PaginatedResponse

	request, err := http.NewRequest(
		http.MethodGet,
		TimeTracking.endpoint,
		nil,
	)

	if err != nil {
		return payload, err
	}

	TimeTracking.client.createQueryUrl(request, params)

	if err = TimeTracking.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (TimeTracking *TimeTrackingsResource) GetTimeTrack(id int) (map[string]interface{}, error) {
	var payload map[string]interface{}

	request, err := http.NewRequest(
		http.MethodGet,
		TimeTracking.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return payload, err
	}

	if err = TimeTracking.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (TimeTracking *TimeTrackingsResource) CreateTimeTrack(body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		TimeTracking.endpoint,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = TimeTracking.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (TimeTracking *TimeTrackingsResource) UpdateTimeTrack(id int, body interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return payload, err
	}

	request, err := http.NewRequest(
		http.MethodPut,
		TimeTracking.endpoint+"/"+strconv.Itoa(id),
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return payload, err
	}

	if err = TimeTracking.client.handle(request, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}

func (TimeTracking *TimeTrackingsResource) DeleteTimeTrack(id int) error {

	request, err := http.NewRequest(
		http.MethodDelete,
		TimeTracking.endpoint+"/"+strconv.Itoa(id),
		nil,
	)

	if err != nil {
		return err
	}

	err = TimeTracking.client.handle(request, nil)

	if err != nil {
		return err
	}

	return nil
}
